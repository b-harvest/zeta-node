package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/evmos/ethermint/x/evm/statedb"
)

// ExtStateDB defines extra methods of statedb to support stateful precompiled contracts
type ExtStateDB interface {
	vm.StateDB
	ExecuteNativeAction(contract common.Address, converter statedb.EventConverter, action func(ctx sdk.Context) error) error
	CacheContext() sdk.Context
}

type Registrable interface {
	RegistryKey() common.Address
}

type BaseContract interface {
	Registrable
}

type baseContract struct {
	address common.Address
}

func NewBaseContract(address common.Address) BaseContract {
	return &baseContract{
		address: address,
	}
}

func (c *baseContract) RegistryKey() common.Address {
	return c.address
}
