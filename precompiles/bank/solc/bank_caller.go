// Copyright Tharsis Labs Ltd.(Evmos)
// SPDX-License-Identifier:ENCL-1.0(https://github.com/evmos/evmos/blob/main/LICENSE)

// TODO: need to refactor due to license issue

package solc

import (
	_ "embed" // embed compiled smart contract
	"encoding/json"

	evmtypes "github.com/evmos/ethermint/x/evm/types"
)

var (
	//go:embed BankCaller.json
	BankCallerJSON []byte

	// BankCallerContract is the compiled contract of BankCaller.sol
	BankCallerContract evmtypes.CompiledContract
)

func init() {
	err := json.Unmarshal(BankCallerJSON, &BankCallerContract)
	if err != nil {
		panic(err)
	}

	if len(BankCallerContract.Bin) == 0 {
		panic("failed to load BankCaller smart contract")
	}
}
