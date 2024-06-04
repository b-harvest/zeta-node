package distribution

import (
	"errors"
	"fmt"
	"math/big"

	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	distributionkeeper "github.com/cosmos/cosmos-sdk/x/distribution/keeper"
	distributiontypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
	"github.com/evmos/ethermint/types"
	evmtypes "github.com/evmos/ethermint/x/evm/types"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/vm"
	ptypes "github.com/zeta-chain/zetacore/precompiles/types"
)

const (
	WithdrawDelegatorRewardsMethodName = "withdrawDelegatorRewards"
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
	ContractAddress     = common.BytesToAddress([]byte{104})
	GasRequiredByMethod = map[[4]byte]uint64{}

	BaseDenom     = "azeta"
	BaseAmount    = sdk.NewInt(100000).Mul(types.PowerReduction)
	BaseCoin      = sdk.NewCoin(BaseDenom, BaseAmount)
	BaseCoins     = sdk.NewCoins(BaseCoin)
	BaseRewardRes = []Coin{
		{
			Denom:  BaseDenom,
			Amount: BaseAmount.BigInt(),
		},
	}
)

func init() {
	ABI, GasRequiredByMethod = initABI()
}

var DistributionModuleMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"delegatorAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"validatorAddress\",\"type\":\"string\"}],\"name\":\"withdrawDelegatorRewards\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"denom\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"struct Coin[]\",\"name\":\"amount\",\"type\":\"tuple[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

func initABI() (abi abi.ABI, gasRequiredByMethod map[[4]byte]uint64) {
	gasRequiredByMethod = map[[4]byte]uint64{}
	if err := abi.UnmarshalJSON([]byte(DistributionModuleMetaData.ABI)); err != nil {
		panic(err)
	}
	for methodName := range abi.Methods {
		var methodID [4]byte
		copy(methodID[:], abi.Methods[methodName].ID[:4])
		switch methodName {
		case WithdrawDelegatorRewardsMethodName:
			gasRequiredByMethod[methodID] = 10
		default:
			gasRequiredByMethod[methodID] = 0
		}
	}
	return abi, gasRequiredByMethod
}

type DistributionContract struct {
	ptypes.BaseContract

	DistributionKeeper distributionkeeper.Keeper
	// temporary keeper due for poc
	BankKeeper  bankkeeper.Keeper
	cdc         codec.Codec
	kvGasConfig storetypes.GasConfig
}

// NewDistributionContract creates the precompiled contract to manage native tokens
func NewDistributionContract(distributionKeeper distributionkeeper.Keeper, bankKeeper bankkeeper.Keeper, cdc codec.Codec, kvGasConfig storetypes.GasConfig) *DistributionContract {
	return &DistributionContract{
		BaseContract:       ptypes.NewBaseContract(ContractAddress),
		DistributionKeeper: distributionKeeper,
		// temporary keeper due for poc
		BankKeeper:  bankKeeper,
		cdc:         cdc,
		kvGasConfig: kvGasConfig,
	}
}

func (dc *DistributionContract) Address() common.Address {
	return ContractAddress
}

func (dc *DistributionContract) Abi() abi.ABI {
	return ABI
}

// RequiredGas calculates the contract gas use
func (dc *DistributionContract) RequiredGas(input []byte) uint64 {
	// base cost to prevent large input size
	baseCost := uint64(len(input)) * dc.kvGasConfig.WriteCostPerByte
	var methodID [4]byte
	copy(methodID[:], input[:4])
	requiredGas, ok := GasRequiredByMethod[methodID]
	if ok {
		return requiredGas + baseCost
	}
	return baseCost
}

// NewMsgWithdrawDelegatorReward creates a new MsgWithdrawDelegatorReward instance.
func NewMsgWithdrawDelegatorReward(args []interface{}) (*distributiontypes.MsgWithdrawDelegatorReward, common.Address, error) {
	if len(args) != 2 {
		return nil, common.Address{}, fmt.Errorf(ptypes.ErrInvalidNumberOfArgs, 2, len(args))
	}

	delegatorAddress, ok := args[0].(common.Address)
	if !ok || delegatorAddress == (common.Address{}) {
		return nil, common.Address{}, fmt.Errorf(ptypes.ErrInvalidDelegator, args[0])
	}

	validatorAddress, _ := args[1].(string)

	msg := &distributiontypes.MsgWithdrawDelegatorReward{
		DelegatorAddress: sdk.AccAddress(delegatorAddress.Bytes()).String(),
		ValidatorAddress: validatorAddress,
	}

	if err := msg.ValidateBasic(); err != nil {
		return nil, common.Address{}, err
	}

	return msg, delegatorAddress, nil
}

// Coin defines a struct that stores all needed information about a coin
// in types native to the EVM.
type Coin struct {
	Denom  string
	Amount *big.Int
}

// NewCoinsResponse converts a response to an array of Coin.
func NewCoinsResponse(amount sdk.Coins) []Coin {
	// Create a new output for each coin and add it to the output array.
	outputs := make([]Coin, len(amount))
	for i, coin := range amount {
		outputs[i] = Coin{
			Denom:  coin.Denom,
			Amount: coin.Amount.BigInt(),
		}
	}
	return outputs
}

func (dc *DistributionContract) WithdrawDelegatorRewards(ctx sdk.Context, origin common.Address, contract *vm.Contract, stateDB vm.StateDB, method *abi.Method, args []interface{}) ([]byte, error) {
	msg, delegatorHexAddr, err := NewMsgWithdrawDelegatorReward(args)
	if err != nil {
		return nil, err
	}

	// If the contract is the delegator, we don't need an origin check
	// Otherwise check if the origin matches the delegator address
	isContractDelegator := contract.CallerAddress == delegatorHexAddr
	if !isContractDelegator && origin != delegatorHexAddr {
		return nil, fmt.Errorf(ptypes.ErrDifferentOrigin, origin.String(), delegatorHexAddr.String())
	}

	delegator, err := sdk.AccAddressFromBech32(msg.DelegatorAddress)
	if err != nil {
		return nil, err
	}

	msgSrv := distributionkeeper.NewMsgServerImpl(dc.DistributionKeeper)
	res, err := msgSrv.WithdrawDelegatorReward(sdk.WrapSDKContext(ctx), msg)
	if err != nil {
		return nil, err
	}

	// ----------------------------------------------------------------------
	// TODO: WARNING!! Artificial mint code for mocking reward generation
	// for PoC purposes only, Do not use it in a production environment.
	if res.Amount.IsZero() {
		err = dc.BankKeeper.MintCoins(
			ctx,
			evmtypes.ModuleName,
			BaseCoins,
		)
		if err != nil {
			return nil, err
		}
		err = dc.BankKeeper.SendCoinsFromModuleToAccount(
			ctx,
			evmtypes.ModuleName,
			delegator,
			BaseCoins,
		)
		if err != nil {
			return nil, err
		}
		res.Amount = BaseCoins
	}
	// ----------------------------------------------------------------------

	return method.Outputs.Pack(NewCoinsResponse(res.Amount))
}

func (dc *DistributionContract) Run(evm *vm.EVM, contract *vm.Contract, readonly bool) ([]byte, error) {
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
	case WithdrawDelegatorRewardsMethodName:
		return dc.WithdrawDelegatorRewards(ctx, evm.Origin, contract, stateDB, method, args)
	// case OtherMethods:
	// ..
	default:
		return nil, errors.New("unknown method")
	}
}
