package producer

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/taikoxyz/taiko-client/bindings"
)

const (
	CircuitsIdx = 0 // Currently we only have one verification contract in protocol.
)

// ProofRequestOptions contains all options that need to be passed to zkEVM rpcd service.
type ProofRequestOptions struct {
	BlockID            *big.Int
	ProverAddress      common.Address
	ProposeBlockTxHash common.Hash
	L1SignalService    common.Address
	L2SignalService    common.Address
	TaikoL2            common.Address
	MetaHash           common.Hash
	BlockHash          common.Hash
	ParentHash         common.Hash
	SignalRoot         common.Hash
	EventL1Hash        common.Hash
	Graffiti           string
	GasUsed            uint64
	ParentGasUsed      uint64
}

type ProofWithHeader struct {
	BlockID *big.Int
	Meta    *bindings.TaikoDataBlockMetadata
	Header  *types.Header
	Proof   []byte
	Degree  uint64
	Opts    *ProofRequestOptions
	Tier    uint16
}

type ProofProducer interface {
	RequestProof(
		ctx context.Context,
		opts *ProofRequestOptions,
		blockID *big.Int,
		meta *bindings.TaikoDataBlockMetadata,
		header *types.Header,
		resultCh chan *ProofWithHeader,
	) error
	Cancellable() bool
	Cancel(ctx context.Context, blockID *big.Int) error
	Tier() uint16
}

func DegreeToCircuitsIdx(degree uint64) (uint16, error) {
	return CircuitsIdx, nil
}
