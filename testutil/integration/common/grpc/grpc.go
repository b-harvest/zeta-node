// Copyright Tharsis Labs Ltd.(Evmos)
// SPDX-License-Identifier:ENCL-1.0(https://github.com/evmos/evmos/blob/main/LICENSE)

package grpc

import (
	sdktypes "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/zeta-chain/zetacore/testutil/integration/common/network"
)

// Handler is an interface that defines the common methods that are used to query
// the network's modules via gRPC.
type Handler interface {
	// Account methods
	GetAccount(address string) (authtypes.AccountI, error)

	// Bank methods
	GetBalance(address sdktypes.AccAddress, denom string) (*banktypes.QueryBalanceResponse, error)
}

var _ Handler = (*IntegrationHandler)(nil)

// IntegrationHandler is a helper struct to query the network's modules
// via gRPC. This is to simulate the behavior of a real user and avoid querying
// the modules directly.
type IntegrationHandler struct {
	network network.Network
}

// NewIntegrationHandler creates a new IntegrationHandler instance.
func NewIntegrationHandler(network network.Network) *IntegrationHandler {
	return &IntegrationHandler{
		network: network,
	}
}
