package keeper_test

import (
	"encoding/hex"
	"errors"
	"math/big"
	"math/rand"
	"testing"

	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/evmos/ethermint/x/evm/statedb"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"github.com/zeta-chain/zetacore/pkg/chains"
	"github.com/zeta-chain/zetacore/pkg/coin"
	keepertest "github.com/zeta-chain/zetacore/testutil/keeper"
	"github.com/zeta-chain/zetacore/testutil/sample"
	"github.com/zeta-chain/zetacore/x/crosschain/keeper"

	"github.com/zeta-chain/zetacore/x/crosschain/types"
	observertypes "github.com/zeta-chain/zetacore/x/observer/types"
)

func setObservers(t *testing.T, k *keeper.Keeper, ctx sdk.Context, zk keepertest.ZetaKeepers) []string {
	validators := k.GetStakingKeeper().GetAllValidators(ctx)

	validatorAddressListFormatted := make([]string, len(validators))
	for i, validator := range validators {
		valAddr, err := sdk.ValAddressFromBech32(validator.OperatorAddress)
		require.NoError(t, err)
		addressTmp, err := sdk.AccAddressFromHexUnsafe(hex.EncodeToString(valAddr.Bytes()))
		require.NoError(t, err)
		validatorAddressListFormatted[i] = addressTmp.String()
	}

	// Add validator to the observer list for voting
	zk.ObserverKeeper.SetObserverSet(ctx, observertypes.ObserverSet{
		ObserverList: validatorAddressListFormatted,
	})
	return validatorAddressListFormatted
}

// TODO: Complete the test cases
// https://github.com/zeta-chain/node/issues/1542
func TestKeeper_VoteOnObservedInboundTx(t *testing.T) {
	t.Run("successfully vote on evm deposit", func(t *testing.T) {
		k, ctx, sdkk, zk := keepertest.CrosschainKeeper(t)
		msgServer := keeper.NewMsgServerImpl(*k)
		validatorList := setObservers(t, k, ctx, zk)
		to, from := int64(1337), int64(101)
		supportedChains := zk.ObserverKeeper.GetSupportedChains(ctx)
		for _, chain := range supportedChains {
			if chains.IsEVMChain(chain.ChainId) {
				from = chain.ChainId
			}
			if chains.IsZetaChain(chain.ChainId) {
				to = chain.ChainId
			}
		}
		zk.ObserverKeeper.SetTSS(ctx, sample.Tss())

		msg := sample.InboundVote(0, from, to)

		err := sdkk.EvmKeeper.SetAccount(ctx, ethcommon.HexToAddress(msg.Receiver), statedb.Account{
			Nonce:    0,
			Balance:  big.NewInt(0),
			CodeHash: crypto.Keccak256(nil),
		})
		require.NoError(t, err)
		for _, validatorAddr := range validatorList {
			msg.Creator = validatorAddr
			_, err := msgServer.VoteOnObservedInboundTx(
				ctx,
				&msg,
			)
			require.NoError(t, err)
		}
		ballot, _, _ := zk.ObserverKeeper.FindBallot(
			ctx,
			msg.Digest(),
			zk.ObserverKeeper.GetSupportedChainFromChainID(ctx, msg.SenderChainId),
			observertypes.ObservationType_InBoundTx,
		)
		require.Equal(t, ballot.BallotStatus, observertypes.BallotStatus_BallotFinalized_SuccessObservation)
		cctx, found := k.GetCrossChainTx(ctx, msg.Digest())
		require.True(t, found)
		require.Equal(t, types.CctxStatus_OutboundMined, cctx.CctxStatus.Status)
		require.Equal(t, cctx.InboundTxParams.TxFinalizationStatus, types.TxFinalizationStatus_Executed)
	})

	t.Run("prevent double event submission", func(t *testing.T) {
		k, ctx, _, zk := keepertest.CrosschainKeeper(t)

		// MsgServer for the crosschain keeper
		msgServer := keeper.NewMsgServerImpl(*k)

		// Convert the validator address into a user address.
		validators := k.GetStakingKeeper().GetAllValidators(ctx)
		validatorAddress := validators[0].OperatorAddress
		valAddr, _ := sdk.ValAddressFromBech32(validatorAddress)
		addresstmp, _ := sdk.AccAddressFromHexUnsafe(hex.EncodeToString(valAddr.Bytes()))
		validatorAddr := addresstmp.String()

		// Add validator to the observer list for voting
		zk.ObserverKeeper.SetObserverSet(ctx, observertypes.ObserverSet{
			ObserverList: []string{validatorAddr},
		})

		// Add tss to the observer keeper
		zk.ObserverKeeper.SetTSS(ctx, sample.Tss())

		// Vote on the FIRST message.
		msg := &types.MsgVoteOnObservedInboundTx{
			Creator:       validatorAddr,
			Sender:        "0x954598965C2aCdA2885B037561526260764095B8",
			SenderChainId: 1337, // ETH
			Receiver:      "0x954598965C2aCdA2885B037561526260764095B8",
			ReceiverChain: 101, // zetachain
			Amount:        sdkmath.NewUintFromString("10000000"),
			Message:       "",
			InBlockHeight: 1,
			GasLimit:      1000000000,
			InTxHash:      "0x7a900ef978743f91f57ca47c6d1a1add75df4d3531da17671e9cf149e1aefe0b",
			CoinType:      0, // zeta
			TxOrigin:      "0x954598965C2aCdA2885B037561526260764095B8",
			Asset:         "",
			EventIndex:    1,
		}
		_, err := msgServer.VoteOnObservedInboundTx(
			ctx,
			msg,
		)
		require.NoError(t, err)

		// Check that the vote passed
		ballot, found := zk.ObserverKeeper.GetBallot(ctx, msg.Digest())
		require.True(t, found)
		require.Equal(t, ballot.BallotStatus, observertypes.BallotStatus_BallotFinalized_SuccessObservation)
		//Perform the SAME event. Except, this time, we resubmit the event.
		msg2 := &types.MsgVoteOnObservedInboundTx{
			Creator:       validatorAddr,
			Sender:        "0x954598965C2aCdA2885B037561526260764095B8",
			SenderChainId: 1337,
			Receiver:      "0x954598965C2aCdA2885B037561526260764095B8",
			ReceiverChain: 101,
			Amount:        sdkmath.NewUintFromString("10000000"),
			Message:       "",
			InBlockHeight: 1,
			GasLimit:      1000000001, // <---- Change here
			InTxHash:      "0x7a900ef978743f91f57ca47c6d1a1add75df4d3531da17671e9cf149e1aefe0b",
			CoinType:      0,
			TxOrigin:      "0x954598965C2aCdA2885B037561526260764095B8",
			Asset:         "",
			EventIndex:    1,
		}

		_, err = msgServer.VoteOnObservedInboundTx(
			ctx,
			msg2,
		)
		require.Error(t, err)
		require.ErrorIs(t, err, types.ErrObservedTxAlreadyFinalized)
		_, found = zk.ObserverKeeper.GetBallot(ctx, msg2.Digest())
		require.False(t, found)
	})

	t.Run("should error if vote on inbound ballot fails", func(t *testing.T) {
		k, ctx, _, _ := keepertest.CrosschainKeeperWithMocks(t, keepertest.CrosschainMockOptions{
			UseObserverMock: true,
		})
		observerMock := keepertest.GetCrosschainObserverMock(t, k)
		observerMock.On("VoteOnInboundBallot", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
			Return(true, false, errors.New("err"))
		msgServer := keeper.NewMsgServerImpl(*k)
		to, from := int64(1337), int64(101)

		msg := sample.InboundVote(0, from, to)
		res, err := msgServer.VoteOnObservedInboundTx(
			ctx,
			&msg,
		)
		require.Error(t, err)
		require.Nil(t, res)
	})

	t.Run("should return if not finalized", func(t *testing.T) {
		k, ctx, _, zk := keepertest.CrosschainKeeper(t)
		msgServer := keeper.NewMsgServerImpl(*k)
		validatorList := setObservers(t, k, ctx, zk)

		// add one more voter to make it not finalized
		r := rand.New(rand.NewSource(42))
		valAddr := sample.ValAddress(r)
		observerSet := append(validatorList, valAddr.String())
		zk.ObserverKeeper.SetObserverSet(ctx, observertypes.ObserverSet{
			ObserverList: observerSet,
		})
		to, from := int64(1337), int64(101)
		supportedChains := zk.ObserverKeeper.GetSupportedChains(ctx)
		for _, chain := range supportedChains {
			if chains.IsEVMChain(chain.ChainId) {
				from = chain.ChainId
			}
			if chains.IsZetaChain(chain.ChainId) {
				to = chain.ChainId
			}
		}
		zk.ObserverKeeper.SetTSS(ctx, sample.Tss())

		msg := sample.InboundVote(0, from, to)
		for _, validatorAddr := range validatorList {
			msg.Creator = validatorAddr
			_, err := msgServer.VoteOnObservedInboundTx(
				ctx,
				&msg,
			)
			require.NoError(t, err)
		}
		ballot, _, _ := zk.ObserverKeeper.FindBallot(
			ctx,
			msg.Digest(),
			zk.ObserverKeeper.GetSupportedChainFromChainID(ctx, msg.SenderChainId),
			observertypes.ObservationType_InBoundTx,
		)
		require.Equal(t, ballot.BallotStatus, observertypes.BallotStatus_BallotInProgress)
		require.Equal(t, ballot.Votes[0], observertypes.VoteType_SuccessObservation)
		require.Equal(t, ballot.Votes[1], observertypes.VoteType_NotYetVoted)
		_, found := k.GetCrossChainTx(ctx, msg.Digest())
		require.False(t, found)
	})

	t.Run("should err if tss not found", func(t *testing.T) {
		k, ctx, _, _ := keepertest.CrosschainKeeperWithMocks(t, keepertest.CrosschainMockOptions{
			UseObserverMock: true,
		})
		observerMock := keepertest.GetCrosschainObserverMock(t, k)
		observerMock.On("VoteOnInboundBallot", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
			Return(true, false, nil)
		observerMock.On("GetTSS", mock.Anything).Return(observertypes.TSS{}, false)
		msgServer := keeper.NewMsgServerImpl(*k)
		to, from := int64(1337), int64(101)

		msg := sample.InboundVote(0, from, to)
		res, err := msgServer.VoteOnObservedInboundTx(
			ctx,
			&msg,
		)
		require.Error(t, err)
		require.Nil(t, res)
	})
}

func TestStatus_ChangeStatus(t *testing.T) {
	tt := []struct {
		Name         string
		Status       types.Status
		NonErrStatus types.CctxStatus
		Msg          string
		IsErr        bool
		ErrStatus    types.CctxStatus
	}{
		{
			Name: "Transition on finalize Inbound",
			Status: types.Status{
				Status:              types.CctxStatus_PendingInbound,
				StatusMessage:       "Getting InTX Votes",
				LastUpdateTimestamp: 0,
			},
			Msg:          "Got super majority and finalized Inbound",
			NonErrStatus: types.CctxStatus_PendingOutbound,
			ErrStatus:    types.CctxStatus_Aborted,
			IsErr:        false,
		},
		{
			Name: "Transition on finalize Inbound Fail",
			Status: types.Status{
				Status:              types.CctxStatus_PendingInbound,
				StatusMessage:       "Getting InTX Votes",
				LastUpdateTimestamp: 0,
			},
			Msg:          "Got super majority and finalized Inbound",
			NonErrStatus: types.CctxStatus_OutboundMined,
			ErrStatus:    types.CctxStatus_Aborted,
			IsErr:        false,
		},
	}
	for _, test := range tt {
		test := test
		t.Run(test.Name, func(t *testing.T) {
			test.Status.ChangeStatus(test.NonErrStatus, test.Msg)
			if test.IsErr {
				require.Equal(t, test.ErrStatus, test.Status.Status)
			} else {
				require.Equal(t, test.NonErrStatus, test.Status.Status)
			}
		})
	}
}

func TestKeeper_SaveInbound(t *testing.T) {
	t.Run("should save the cctx", func(t *testing.T) {
		k, ctx, _, zk := keepertest.CrosschainKeeper(t)
		zk.ObserverKeeper.SetTSS(ctx, sample.Tss())
		receiver := sample.EthAddress()
		amount := big.NewInt(42)
		senderChain := getValidEthChain()
		cctx := GetERC20Cctx(t, receiver, *senderChain, "", amount)
		eventIndex := sample.Uint64InRange(1, 100)
		k.SaveInbound(ctx, cctx, eventIndex)
		require.Equal(t, types.TxFinalizationStatus_Executed, cctx.InboundTxParams.TxFinalizationStatus)
		require.True(t, k.IsFinalizedInbound(ctx, cctx.GetInboundTxParams().InboundTxObservedHash, cctx.GetInboundTxParams().SenderChainId, eventIndex))
		_, found := k.GetCrossChainTx(ctx, cctx.Index)
		require.True(t, found)
	})

	t.Run("should save the cctx and remove tracker", func(t *testing.T) {
		k, ctx, _, zk := keepertest.CrosschainKeeper(t)
		receiver := sample.EthAddress()
		amount := big.NewInt(42)
		senderChain := getValidEthChain()
		cctx := GetERC20Cctx(t, receiver, *senderChain, "", amount)
		hash := sample.Hash()
		cctx.InboundTxParams.InboundTxObservedHash = hash.String()
		k.SetInTxTracker(ctx, types.InTxTracker{
			ChainId:  senderChain.ChainId,
			TxHash:   hash.String(),
			CoinType: 0,
		})
		eventIndex := sample.Uint64InRange(1, 100)
		zk.ObserverKeeper.SetTSS(ctx, sample.Tss())

		k.SaveInbound(ctx, cctx, eventIndex)
		require.Equal(t, types.TxFinalizationStatus_Executed, cctx.InboundTxParams.TxFinalizationStatus)
		require.True(t, k.IsFinalizedInbound(ctx, cctx.GetInboundTxParams().InboundTxObservedHash, cctx.GetInboundTxParams().SenderChainId, eventIndex))
		_, found := k.GetCrossChainTx(ctx, cctx.Index)
		require.True(t, found)
		_, found = k.GetInTxTracker(ctx, senderChain.ChainId, hash.String())
		require.False(t, found)
	})
}

// GetERC20Cctx returns a sample CrossChainTx with ERC20 params. This is used for testing Inbound and Outbound voting transactions
func GetERC20Cctx(t *testing.T, receiver ethcommon.Address, senderChain chains.Chain, asset string, amount *big.Int) *types.CrossChainTx {
	r := sample.Rand()
	cctx := &types.CrossChainTx{
		Creator:          sample.AccAddress(),
		Index:            sample.ZetaIndex(t),
		ZetaFees:         sample.UintInRange(0, 100),
		RelayedMessage:   "",
		CctxStatus:       &types.Status{Status: types.CctxStatus_PendingInbound},
		InboundTxParams:  sample.InboundTxParams(r),
		OutboundTxParams: []*types.OutboundTxParams{sample.OutboundTxParams(r)},
	}

	cctx.GetInboundTxParams().Amount = sdkmath.NewUintFromBigInt(amount)
	cctx.GetInboundTxParams().SenderChainId = senderChain.ChainId
	cctx.GetInboundTxParams().InboundTxObservedHash = sample.Hash().String()
	cctx.GetInboundTxParams().InboundTxBallotIndex = sample.ZetaIndex(t)

	cctx.GetCurrentOutTxParam().ReceiverChainId = senderChain.ChainId
	cctx.GetCurrentOutTxParam().Receiver = receiver.String()
	cctx.GetCurrentOutTxParam().OutboundTxHash = sample.Hash().String()
	cctx.GetCurrentOutTxParam().OutboundTxBallotIndex = sample.ZetaIndex(t)

	cctx.InboundTxParams.CoinType = coin.CoinType_ERC20
	for _, outboundTxParam := range cctx.OutboundTxParams {
		outboundTxParam.CoinType = coin.CoinType_ERC20
	}

	cctx.GetInboundTxParams().Asset = asset
	cctx.GetInboundTxParams().Sender = sample.EthAddress().String()
	cctx.GetCurrentOutTxParam().OutboundTxTssNonce = 42
	cctx.GetCurrentOutTxParam().OutboundTxGasUsed = 100
	cctx.GetCurrentOutTxParam().OutboundTxEffectiveGasLimit = 100
	return cctx
}
