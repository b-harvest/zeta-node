// Copyright Tharsis Labs Ltd.(Evmos)
// SPDX-License-Identifier:ENCL-1.0(https://github.com/evmos/evmos/blob/main/LICENSE)

package solc

import (
	_ "embed" // embed compiled smart contract
	"encoding/json"
	evmtypes "github.com/evmos/ethermint/x/evm/types"
)

var (
	//go:embed RegularCaller.json
	RegularCallerJSON []byte

	// RegularCallerContract is the compiled contract of RegularCaller.sol
	RegularCallerContract evmtypes.CompiledContract
)

func init() {
	err := json.Unmarshal(RegularCallerJSON, &RegularCallerContract)
	if err != nil {
		panic(err)
	}

	if len(RegularCallerContract.Bin) == 0 {
		panic("failed to load RegularCaller smart contract")
	}
}
