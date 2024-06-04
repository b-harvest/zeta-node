package types

import (
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

// Balance contains the amount for a corresponding ERC-20 contract address
type Balance struct {
	ContractAddress common.Address
	Amount          *big.Int
}

func BytesToBigInt(data []byte) *big.Int {
	return big.NewInt(0).SetBytes(data[:])
}
