package crosschain_test

import (
	"fmt"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	evmtypes "github.com/evmos/ethermint/x/evm/types"
	"github.com/zeta-chain/zetacore/testutil/integration/zeta/factory"
	//nolint:revive // dot imports are fine for Ginkgo
)

// callType constants to differentiate between direct calls and calls through a contract.
const (
	directCall = iota + 1
	contractCall
)

// ContractData is a helper struct to hold the addresses and ABIs for the
// different contract instances that are subject to testing here.
type ContractData struct {
	ownerPriv cryptotypes.PrivKey

	contractAddr   common.Address
	contractABI    abi.ABI
	precompileAddr common.Address
	precompileABI  abi.ABI
}

func (cd ContractData) String() string {
	return fmt.Sprintf(
		"ContractData{\n\townerPriv: %x,\n\tprecompileAddr: %s,\n\tprecompileABI: %s,\n\tcontractAddr: %s,\n\tcontractABI: %s\n}",
		cd.ownerPriv,
		cd.precompileAddr.Hex(),
		cd.precompileABI.Methods,
		cd.contractAddr.Hex(),
		cd.contractABI.Methods,
	)
}

// getCallArgs is a helper function to return the correct call arguments for a given call type.
// In case of a direct call to the precompile, the precompile's ABI is used. Otherwise a caller contract is used.
func getTxAndCallArgs(
	callType int,
	contractData ContractData,
	methodName string,
	args ...interface{},
) (evmtypes.TransactionArgs, factory.CallArgs) {
	txArgs := evmtypes.TransactionArgs{}
	callArgs := factory.CallArgs{}

	switch callType {
	case directCall:
		txArgs.To = &contractData.precompileAddr
		callArgs.ContractABI = contractData.precompileABI
	case contractCall:
		txArgs.To = &contractData.contractAddr
		callArgs.ContractABI = contractData.contractABI
	}

	callArgs.MethodName = methodName
	callArgs.Args = args

	return txArgs, callArgs
}
