// Copyright Tharsis Labs Ltd.(Evmos)
// SPDX-License-Identifier:ENCL-1.0(https://github.com/evmos/evmos/blob/main/LICENSE)
package factory

import (
	"github.com/ethereum/go-ethereum/common"
	evmtypes "github.com/evmos/ethermint/x/evm/types"

	errorsmod "cosmossdk.io/errors"
	amino "github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	testutiltypes "github.com/cosmos/cosmos-sdk/types/module/testutil"
	authtx "github.com/cosmos/cosmos-sdk/x/auth/tx"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	etherminttypes "github.com/evmos/ethermint/types"
	"github.com/zeta-chain/zetacore/testutil/tx"
)

// NewTx returns a reference to a new Ethereum transaction message.
//func NewTx(
//	tx *evmtypes.TransactionArgs,
//) *evmtypes.MsgEthereumTx {
//	return newMsgEthereumTx(tx)
//}

//
//func newMsgEthereumTx(
//	tx *evmtypes.TransactionArgs,
//) *evmtypes.MsgEthereumTx {
//	var (
//		cid, amt, gp *sdkmath.Int
//		toAddr       string
//		txData       evmtypes.TxData
//	)
//
//	if tx.To != nil {
//		toAddr = tx.To.Hex()
//	}
//
//	if tx.Amount != nil {
//		amountInt := sdkmath.NewIntFromBigInt(tx.Amount)
//		amt = &amountInt
//	}
//
//	if tx.ChainID != nil {
//		chainIDInt := sdkmath.NewIntFromBigInt(tx.ChainID)
//		cid = &chainIDInt
//	}
//
//	if tx.GasPrice != nil {
//		gasPriceInt := sdkmath.NewIntFromBigInt(tx.GasPrice)
//		gp = &gasPriceInt
//	}
//
//	switch {
//	case tx.GasFeeCap != nil:
//		gtc := sdkmath.NewIntFromBigInt(tx.GasTipCap)
//		gfc := sdkmath.NewIntFromBigInt(tx.GasFeeCap)
//
//		txData = &DynamicFeeTx{
//			ChainID:   cid,
//			Amount:    amt,
//			To:        toAddr,
//			GasTipCap: &gtc,
//			GasFeeCap: &gfc,
//			Nonce:     tx.Nonce,
//			GasLimit:  tx.GasLimit,
//			Data:      tx.Input,
//			Accesses:  NewAccessList(tx.Accesses),
//		}
//	case tx.Accesses != nil:
//		txData = &AccessListTx{
//			ChainID:  cid,
//			Nonce:    tx.Nonce,
//			To:       toAddr,
//			Amount:   amt,
//			GasLimit: tx.GasLimit,
//			GasPrice: gp,
//			Data:     tx.Input,
//			Accesses: NewAccessList(tx.Accesses),
//		}
//	default:
//		txData = &LegacyTx{
//			To:       toAddr,
//			Amount:   amt,
//			GasPrice: gp,
//			Nonce:    tx.Nonce,
//			GasLimit: tx.GasLimit,
//			Data:     tx.Input,
//		}
//	}
//
//	dataAny, err := PackTxData(txData)
//	if err != nil {
//		panic(err)
//	}
//
//	msg := MsgEthereumTx{Data: dataAny}
//	msg.Hash = msg.AsTransaction().Hash().Hex()
//	return &msg
//}

// buildMsgEthereumTx builds an Ethereum transaction from the given arguments and populates the From field.
func buildMsgEthereumTx(txArgs evmtypes.TransactionArgs, fromAddr common.Address) evmtypes.MsgEthereumTx {
	//msgEthereumTx := evmtypes.NewTx(&txArgs)
	msgEthereumTx := txArgs.ToTransaction()
	msgEthereumTx.From = fromAddr.String()
	return *msgEthereumTx
}

// signMsgEthereumTx signs a MsgEthereumTx with the provided private key and chainID.
func signMsgEthereumTx(msgEthereumTx evmtypes.MsgEthereumTx, privKey cryptotypes.PrivKey, chainID string) (evmtypes.MsgEthereumTx, error) {
	ethChainID, err := etherminttypes.ParseChainID(chainID)
	if err != nil {
		return evmtypes.MsgEthereumTx{}, errorsmod.Wrapf(err, "failed to parse chainID: %v", chainID)
	}

	signer := ethtypes.LatestSignerForChainID(ethChainID)
	err = msgEthereumTx.Sign(signer, tx.NewSigner(privKey))
	if err != nil {
		return evmtypes.MsgEthereumTx{}, errorsmod.Wrap(err, "failed to sign transaction")
	}

	// Validate the transaction to avoid unrealistic behavior
	if err = msgEthereumTx.ValidateBasic(); err != nil {
		return evmtypes.MsgEthereumTx{}, errorsmod.Wrap(err, "failed to validate transaction")
	}
	return msgEthereumTx, nil
}

// makeConfig creates an EncodingConfig for testing
func makeConfig(mb module.BasicManager) testutiltypes.TestEncodingConfig {
	cdc := amino.NewLegacyAmino()
	interfaceRegistry := codectypes.NewInterfaceRegistry()
	codec := amino.NewProtoCodec(interfaceRegistry)

	encodingConfig := testutiltypes.TestEncodingConfig{
		InterfaceRegistry: interfaceRegistry,
		Codec:             codec,
		TxConfig:          authtx.NewTxConfig(codec, authtx.DefaultSignModes),
		Amino:             cdc,
	}

	//enccodec.RegisterLegacyAminoCodec(encodingConfig.Amino)
	mb.RegisterLegacyAminoCodec(encodingConfig.Amino)
	//enccodec.RegisterInterfaces(encodingConfig.InterfaceRegistry)
	mb.RegisterInterfaces(encodingConfig.InterfaceRegistry)
	return encodingConfig
}
