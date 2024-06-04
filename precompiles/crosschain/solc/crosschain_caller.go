// Copyright Tharsis Labs Ltd.(Evmos)
// SPDX-License-Identifier:ENCL-1.0(https://github.com/evmos/evmos/blob/main/LICENSE)

package solc

import (
	_ "embed" // embed compiled smart contract
	"encoding/json"
	evmtypes "github.com/evmos/ethermint/x/evm/types"
)

var (
	//go:embed CrossChainCaller.json
	CrossChainCallerJSON []byte

	// CrossChainCallerContract is the compiled contract of CrossChainCaller.sol
	CrossChainCallerContract evmtypes.CompiledContract
)

func init() {
	err := json.Unmarshal(CrossChainCallerJSON, &CrossChainCallerContract)
	if err != nil {
		panic(err)
	}

	if len(CrossChainCallerContract.Bin) == 0 {
		panic("failed to load CrossChainCaller smart contract")
	}
}
