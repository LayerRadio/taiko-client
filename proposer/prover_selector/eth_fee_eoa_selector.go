package selector

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"math/rand"
	"net/url"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/log"
	"github.com/go-resty/resty/v2"
	"github.com/taikoxyz/taiko-client/bindings"
	"github.com/taikoxyz/taiko-client/bindings/encoding"
	"github.com/taikoxyz/taiko-client/pkg/rpc"
	"github.com/taikoxyz/taiko-client/prover/server"
)

var (
	httpScheme              = "http"
	httpsScheme             = "https"
	errEmptyProverEndpoints = errors.New("empty prover endpoints")
	errUnableToFindProver   = errors.New("unable to find prover")
)

// ETHFeeEOASelector is a prover selector implementation which use ETHs as prover fee and
// all provers selected must be EOA accounts.
type ETHFeeEOASelector struct {
	protocolConfigs               *bindings.TaikoDataConfig
	rpc                           *rpc.Client
	taikoL1Address                common.Address
	tiersFee                      []encoding.TierFee
	tierFeePriceBump              *big.Int
	proverEndpoints               []*url.URL
	maxTierFeePriceBumpIterations uint64
	proposalExpiry                time.Duration
	requestTimeout                time.Duration
}

// NewETHFeeEOASelector creates a new ETHFeeEOASelector instance.
func NewETHFeeEOASelector(
	protocolConfigs *bindings.TaikoDataConfig,
	rpc *rpc.Client,
	taikoL1Address common.Address,
	tiersFee []encoding.TierFee,
	tierFeePriceBump *big.Int,
	proverEndpoints []*url.URL,
	maxTierFeePriceBumpIterations uint64,
	proposalExpiry time.Duration,
	requestTimeout time.Duration,
) (*ETHFeeEOASelector, error) {
	if len(proverEndpoints) == 0 {
		return nil, errEmptyProverEndpoints
	}

	for _, endpoint := range proverEndpoints {
		if endpoint.Scheme != httpScheme && endpoint.Scheme != httpsScheme {
			return nil, fmt.Errorf("invalid prover endpoint %s", endpoint)
		}
	}

	return &ETHFeeEOASelector{
		protocolConfigs,
		rpc,
		taikoL1Address,
		tiersFee,
		tierFeePriceBump,
		proverEndpoints,
		maxTierFeePriceBumpIterations,
		proposalExpiry,
		requestTimeout,
	}, nil
}

// ProverEndpoints returns all registered prover endpoints.
func (s *ETHFeeEOASelector) ProverEndpoints() []*url.URL { return s.proverEndpoints }

// AssignProver tries to pick a prover through the registered prover endpoints.
func (s *ETHFeeEOASelector) AssignProver(
	ctx context.Context,
	tierFees []encoding.TierFee,
	txListHash common.Hash,
) ([]byte, *big.Int, error) {
	guardianProverAddress, err := s.rpc.TaikoL1.Resolve0(
		&bind.CallOpts{Context: ctx},
		rpc.StringToBytes32("guardian"),
		true,
	)
	if err != nil {
		return nil, nil, err
	}

	var (
		expiry       = uint64(time.Now().Add(s.proposalExpiry).Unix())
		fees         = make([]encoding.TierFee, len(tierFees))
		big100       = new(big.Int).SetUint64(uint64(100))
		maxProverFee = common.Big0
	)
	copy(fees, tierFees)

	// Iterate over each configured endpoint, and see if someone wants to accept this block.
	// If it is denied, we continue on to the next endpoint.
	// If we do not find a prover, we can increase the fee up to a point, or give up.
	for i := 0; i < int(s.maxTierFeePriceBumpIterations); i++ {
		// Bump tier fee on each failed loop
		cumulativeBumpPercent := new(big.Int).Mul(s.tierFeePriceBump, new(big.Int).SetUint64(uint64(i)))
		for idx := range fees {
			if i > 0 {
				fee := new(big.Int).Mul(fees[idx].Fee, cumulativeBumpPercent)
				fees[idx].Fee = fees[idx].Fee.Add(fees[idx].Fee, fee.Div(fee, big100))
			}
			if fees[idx].Fee.Cmp(maxProverFee) > 0 {
				maxProverFee = fees[idx].Fee
			}
		}

		for _, endpoint := range s.shuffleProverEndpoints() {
			encodedAssignment, proverAddress, err := assignProver(
				ctx,
				endpoint,
				expiry,
				tierFees,
				txListHash,
				s.requestTimeout,
				guardianProverAddress,
			)
			if err != nil {
				log.Warn("Failed to assign prover", "endpoint", endpoint, "error", err)
				continue
			}

			ok, err := rpc.CheckProverBalance(ctx, s.rpc, proverAddress, s.taikoL1Address, s.protocolConfigs.LivenessBond)
			if err != nil {
				log.Warn("Failed to check prover balance", "endpoint", endpoint, "error", err)
				continue
			}
			if !ok {
				continue
			}

			return encodedAssignment, maxProverFee, nil
		}
	}

	return nil, nil, errUnableToFindProver
}

// shuffleProverEndpoints shuffles the current selector's prover endpoints.
func (s *ETHFeeEOASelector) shuffleProverEndpoints() []*url.URL {
	rand.Shuffle(len(s.proverEndpoints), func(i, j int) {
		s.proverEndpoints[i], s.proverEndpoints[j] = s.proverEndpoints[j], s.proverEndpoints[i]
	})
	return s.proverEndpoints
}

// assignProver tries to assign a proof generation task to the given prover by HTTP API.
func assignProver(
	ctx context.Context,
	endpoint *url.URL,
	expiry uint64,
	tierFees []encoding.TierFee,
	txListHash common.Hash,
	timeout time.Duration,
	guardianProverAddress common.Address,
) ([]byte, common.Address, error) {
	log.Info(
		"Attempting to assign prover",
		"endpoint", endpoint,
		"expiry", expiry,
		"txListHash", txListHash,
	)

	// Send the HTTP request
	var (
		client  = resty.New()
		reqBody = &server.CreateAssignmentRequestBody{
			FeeToken:   (common.Address{}),
			TierFees:   tierFees,
			Expiry:     expiry,
			TxListHash: txListHash,
		}
		result = server.ProposeBlockResponse{}
	)
	requestUrl, err := url.JoinPath(endpoint.String(), "/assignment")
	if err != nil {
		return nil, common.Address{}, err
	}

	ctxTimeout, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	resp, err := client.R().
		SetContext(ctxTimeout).
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(reqBody).
		SetResult(&result).
		Post(requestUrl)
	if err != nil {
		return nil, common.Address{}, err
	}
	if !resp.IsSuccess() {
		return nil, common.Address{}, fmt.Errorf("unsuccessful response %d", resp.StatusCode())
	}

	// Ensure prover in response is the same as the one recovered
	// from the signature
	payload, err := encoding.EncodeProverAssignmentPayload(txListHash, common.Address{}, expiry, tierFees)
	if err != nil {
		return nil, common.Address{}, err
	}

	pubKey, err := crypto.SigToPub(crypto.Keccak256Hash(payload).Bytes(), result.SignedPayload)
	if err != nil {
		return nil, common.Address{}, err
	}

	if crypto.PubkeyToAddress(*pubKey).Hex() != result.Prover.Hex() {
		return nil, common.Address{}, fmt.Errorf(
			"assigned prover signature did not recover to provided prover address %s != %s",
			crypto.PubkeyToAddress(*pubKey).Hex(),
			result.Prover.Hex(),
		)
	}

	// Convert signature to one solidity can recover by adding 27 to 65th byte
	result.SignedPayload[64] = uint8(uint(result.SignedPayload[64])) + 27

	encoded, err := encoding.EncodeProverAssignment(&encoding.ProverAssignment{
		Prover:    result.Prover,
		FeeToken:  common.Address{},
		TierFees:  tierFees,
		Expiry:    reqBody.Expiry,
		Signature: result.SignedPayload,
	})
	if err != nil {
		return nil, common.Address{}, err
	}

	log.Info(
		"Prover assigned",
		"address", result.Prover,
		"endpoint", endpoint,
		"tierFees", tierFees,
		"expiry", expiry,
	)

	return encoded, result.Prover, nil
}
