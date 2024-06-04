package types

import (
	"bytes"
	"embed"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

// TODO: refactor to License
// LoadABI read the ABI file described by the path and parse it as JSON.
func LoadABI(fs embed.FS, path string) (abi.ABI, error) {
	abiBz, err := fs.ReadFile(path)
	if err != nil {
		return abi.ABI{}, fmt.Errorf("error loading the ABI %s", err)
	}

	newAbi, err := abi.JSON(bytes.NewReader(abiBz))
	if err != nil {
		return abi.ABI{}, fmt.Errorf(ErrInvalidABI, err)
	}
	return newAbi, nil
}

// Balance contains the amount for a corresponding ERC-20 contract address
type Balance struct {
	ContractAddress common.Address
	Amount          *big.Int
}
