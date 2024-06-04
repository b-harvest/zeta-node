package crosschain

import (
	"errors"
	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/vm"
	ptypes "github.com/zeta-chain/zetacore/precompiles/types"
)

const (
	GasPriceMethodName = "gasPrice"
)

type GasPriceRes struct {
	Creator     string
	Index       string
	ChainId     int64
	Signers     []string
	BlockNums   []uint64
	Prices      []uint64
	MedianIndex uint64
	Found       bool
}

var (
	ABI                 abi.ABI
	ContractAddress     = common.BytesToAddress([]byte{200})
	GasRequiredByMethod = map[[4]byte]uint64{}
)

func init() {
	ABI, GasRequiredByMethod = initABI()
}

var CrossChainModuleMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"int64\",\"name\":\"chainID\",\"type\":\"int64\"}],\"name\":\"gasPrice\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"creator\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"index\",\"type\":\"string\"},{\"internalType\":\"int64\",\"name\":\"chainId\",\"type\":\"int64\"},{\"internalType\":\"string[]\",\"name\":\"signers\",\"type\":\"string[]\"},{\"internalType\":\"uint64[]\",\"name\":\"blockNums\",\"type\":\"uint64[]\"},{\"internalType\":\"uint64[]\",\"name\":\"prices\",\"type\":\"uint64[]\"},{\"internalType\":\"uint64\",\"name\":\"medianIndex\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"found\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

func initABI() (abi abi.ABI, gasRequiredByMethod map[[4]byte]uint64) {
	gasRequiredByMethod = map[[4]byte]uint64{}
	if err := abi.UnmarshalJSON([]byte(CrossChainModuleMetaData.ABI)); err != nil {
		panic(err)
	}
	for methodName := range abi.Methods {
		var methodID [4]byte
		copy(methodID[:], abi.Methods[methodName].ID[:4])
		switch methodName {
		case GasPriceMethodName:
			gasRequiredByMethod[methodID] = 10
		default:
			gasRequiredByMethod[methodID] = 0
		}
	}
	return abi, gasRequiredByMethod
}

type CrossChainContract struct {
	ptypes.BaseContract

	CrossChainKeeper crosschainKeeper
	cdc              codec.Codec
	kvGasConfig      storetypes.GasConfig
}

// NewCrossChainContract creates the precompiled contract to manage native tokens
func NewCrossChainContract(crossChainKeeper crosschainKeeper, cdc codec.Codec, kvGasConfig storetypes.GasConfig) *CrossChainContract {
	return &CrossChainContract{
		BaseContract:     ptypes.NewBaseContract(ContractAddress),
		CrossChainKeeper: crossChainKeeper,
		cdc:              cdc,
		kvGasConfig:      kvGasConfig,
	}
}

func (bc *CrossChainContract) Address() common.Address {
	return ContractAddress
}

func (bc *CrossChainContract) Abi() abi.ABI {
	return ABI
}

// RequiredGas calculates the contract gas use
func (bc *CrossChainContract) RequiredGas(input []byte) uint64 {
	// base cost to prevent large input size
	baseCost := uint64(len(input)) * bc.kvGasConfig.WriteCostPerByte
	var methodID [4]byte
	copy(methodID[:], input[:4])
	requiredGas, ok := GasRequiredByMethod[methodID]
	if ok {
		return requiredGas + baseCost
	}
	return baseCost
}

func (bc *CrossChainContract) GasPrice(ctx sdk.Context, method *abi.Method, args []interface{}) ([]byte, error) {
	chainId := args[0].(int64)

	gp, found := bc.CrossChainKeeper.GetGasPrice(ctx, chainId)

	return method.Outputs.Pack(
		gp.Creator,
		gp.Index,
		gp.ChainId,
		gp.Signers,
		gp.BlockNums,
		gp.Prices,
		gp.MedianIndex,
		found,
	)
}

func (bc *CrossChainContract) Run(evm *vm.EVM, contract *vm.Contract, readonly bool) ([]byte, error) {
	// parse input
	methodID := contract.Input[:4]
	method, err := ABI.MethodById(methodID)
	if err != nil {
		return nil, err
	}
	args, err := method.Inputs.Unpack(contract.Input[4:])
	if err != nil {
		return nil, errors.New("fail to unpack input arguments")
	}

	stateDB := evm.StateDB.(ptypes.ExtStateDB)
	ctx := stateDB.CacheContext()

	switch method.Name {
	case GasPriceMethodName:
		return bc.GasPrice(ctx, method, args)
	default:
		return nil, errors.New("unknown method")
	}
}
