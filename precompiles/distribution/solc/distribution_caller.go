// Copyright Tharsis Labs Ltd.(Evmos)
// SPDX-License-Identifier:ENCL-1.0(https://github.com/evmos/evmos/blob/main/LICENSE)

package solc

import (
	_ "embed" // embed compiled smart contract
	"encoding/json"
	evmtypes "github.com/evmos/ethermint/x/evm/types"
)

var (
	//go:embed DistributionCaller.json
	DistributionCallerJSON []byte

	// DistributionCallerContract is the compiled contract of DistributionCaller.sol
	DistributionCallerContract evmtypes.CompiledContract
)

func init() {
	err := json.Unmarshal(DistributionCallerJSON, &DistributionCallerContract)
	if err != nil {
		panic(err)
	}

	if len(DistributionCallerContract.Bin) == 0 {
		panic("failed to load DistributionCaller smart contract")
	}
}
