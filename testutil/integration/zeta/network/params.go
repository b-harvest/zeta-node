// Copyright Tharsis Labs Ltd.(Evmos)
// SPDX-License-Identifier:ENCL-1.0(https://github.com/evmos/evmos/blob/main/LICENSE)

package network

import (
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types/v1"
	evmtypes "github.com/evmos/ethermint/x/evm/types"
)

func (n *IntegrationNetwork) UpdateEvmParams(params evmtypes.Params) error {
	return n.app.EvmKeeper.SetParams(n.ctx, params)
}

func (n *IntegrationNetwork) UpdateGovParams(params govtypes.Params) error {
	return n.app.GovKeeper.SetParams(n.ctx, params)
}
