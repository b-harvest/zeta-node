// Copyright Tharsis Labs Ltd.(Evmos)
// SPDX-License-Identifier:ENCL-1.0(https://github.com/evmos/evmos/blob/main/LICENSE)

package factory

import (
	"github.com/ethereum/go-ethereum/accounts/abi"
	evmtypes "github.com/evmos/ethermint/x/evm/types"
)

// CallArgs is a struct to define all relevant data to call a smart contract.
type CallArgs struct {
	// ContractABI is the ABI of the contract to call.
	ContractABI abi.ABI
	// MethodName is the name of the method to call.
	MethodName string
	// Args are the arguments to pass to the method.
	Args []interface{}
}

// ContractDeploymentData is a struct to define all relevant data to deploy a smart contract.
type ContractDeploymentData struct {
	// Contract is the compiled contract to deploy.
	Contract evmtypes.CompiledContract
	// ConstructorArgs are the arguments to pass to the constructor.
	ConstructorArgs []interface{}
}
