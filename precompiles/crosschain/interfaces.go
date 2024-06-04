package crosschain

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/zeta-chain/zetacore/x/crosschain/types"
)

// crosschainKeeper is an interface to prevent cyclic dependency
type crosschainKeeper interface {
	GetGasPrice(ctx sdk.Context, chainID int64) (val types.GasPrice, found bool)
}
