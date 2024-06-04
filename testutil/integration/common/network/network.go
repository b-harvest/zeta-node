// Copyright Tharsis Labs Ltd.(Evmos)
// SPDX-License-Identifier:ENCL-1.0(https://github.com/evmos/evmos/blob/main/LICENSE)

package network

import (
	"time"

	sdkmath "cosmossdk.io/math"
	abcitypes "github.com/cometbft/cometbft/abci/types"
	sdktypes "github.com/cosmos/cosmos-sdk/types"
	txtypes "github.com/cosmos/cosmos-sdk/types/tx"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/cosmos/cosmos-sdk/x/authz"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
)

// Network is the interface that wraps the common methods to interact with integration test network.
//
// It was designed to avoid users to access module's keepers directly and force integration tests
// to be closer to the real user's behavior.
type Network interface {
	GetContext() sdktypes.Context
	GetChainID() string
	GetDenom() string
	GetValidators() []stakingtypes.Validator

	NextBlock() error
	NextBlockAfter(duration time.Duration) error

	// Clients
	GetAuthClient() authtypes.QueryClient
	GetAuthzClient() authz.QueryClient
	GetBankClient() banktypes.QueryClient
	GetStakingClient() stakingtypes.QueryClient

	BroadcastTxSync(txBytes []byte) (abcitypes.ResponseDeliverTx, error)
	Simulate(txBytes []byte) (*txtypes.SimulateResponse, error)

	// FundAccount funds the given account with the given amount.
	FundAccount(address sdktypes.AccAddress, amount sdktypes.Coins) error
	// FundAccountWithBaseDenom funds the given account with the given amount of the network's
	// base denomination.
	FundAccountWithBaseDenom(address sdktypes.AccAddress, amount sdkmath.Int) error
}
