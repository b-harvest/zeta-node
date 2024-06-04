package distribution_test

import (
	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/vm"
	evmtypes "github.com/evmos/ethermint/x/evm/types"
	"github.com/zeta-chain/zetacore/app"
	"github.com/zeta-chain/zetacore/precompiles/distribution"
	"github.com/zeta-chain/zetacore/precompiles/distribution/solc"
	"github.com/zeta-chain/zetacore/precompiles/testutil"
	"github.com/zeta-chain/zetacore/testutil/integration/zeta/factory"
	"github.com/zeta-chain/zetacore/testutil/integration/zeta/grpc"
	"github.com/zeta-chain/zetacore/testutil/integration/zeta/network"
	"math/big"
	"testing"

	"github.com/zeta-chain/zetacore/testutil/integration/zeta/keyring"

	//nolint:revive // dot imports are fine for Ginkgo
	. "github.com/onsi/ginkgo/v2"
	//nolint:revive // dot imports are fine for Ginkgo
	. "github.com/onsi/gomega"
)

var is *IntegrationTestSuite

// IntegrationTestSuite is the implementation of the TestSuite interface for Distribution precompile
// unit testis.
type IntegrationTestSuite struct {
	bondDenom  string
	addr       common.Address
	addr2      common.Address
	delegation types.Delegation
	validator  types.Validator

	ctx sdk.Context

	// tokenDenom is the specific token denomination used in testing the Distribution precompile.
	// This denomination is used to instantiate the precompile.
	network *network.UnitTestNetwork
	factory factory.TxFactory
	keyring keyring.Keyring

	precompile          *distribution.DistributionContract
	precompiledContract vm.PrecompiledContract
	abi                 abi.ABI

	kvGasConfig storetypes.GasConfig
	cdc         codec.Codec
}

func (is *IntegrationTestSuite) SetupTest() {
	is.kvGasConfig = storetypes.TransientGasConfig()
	encodingConfig := app.MakeEncodingConfig()
	is.cdc = encodingConfig.Codec

	keyring := keyring.New(2)
	integrationNetwork := network.NewUnitTestNetwork(
		network.WithPreFundedAccounts(keyring.GetAllAccAddrs()...),
	)
	grpcHandler := grpc.NewIntegrationHandler(integrationNetwork)
	txFactory := factory.New(integrationNetwork, grpcHandler)

	is.ctx = integrationNetwork.GetContext()
	evmGenesis := *evmtypes.DefaultGenesisState()
	evmGenesis.Params.EvmDenom = distribution.BaseDenom
	integrationNetwork.UpdateEvmParams(evmGenesis.Params)
	Expect(distribution.BaseDenom).ToNot(BeEmpty(), "bond denom cannot be empty")

	is.bondDenom = distribution.BaseDenom
	is.factory = txFactory
	is.keyring = keyring
	is.network = integrationNetwork

	is.setupDistributionPrecompile()

	dels := is.network.App.StakingKeeper.GetAllDelegations(is.ctx)
	is.delegation = dels[0]
	is.addr = common.BytesToAddress(is.delegation.GetDelegatorAddr().Bytes())
	is.addr2 = common.BytesToAddress(sdk.MustAccAddressFromBech32("zeta1syavy2npfyt9tcncdtsdzf7kny9lh777heefxk").Bytes())

	validator, found := is.network.App.StakingKeeper.GetValidator(is.ctx, is.delegation.GetValidatorAddr())
	Expect(found).ToNot(BeFalse())
	is.validator = validator
}

func (is *IntegrationTestSuite) setupDistributionPrecompile() {
	pcc := distribution.NewDistributionContract(
		is.network.App.DistrKeeper,
		is.network.App.BankKeeper,
		is.cdc,
		is.kvGasConfig,
	)

	is.precompile = pcc
	is.abi = is.precompile.Abi()
}

// For a concise PoC, reward is minted and sent in WithdrawDelegatorRewards instead of AllocateRewards,
// separately from msg processing. This can be replaced in the future by implementing AllocateRewards as needed.
func (is *IntegrationTestSuite) AllocateRewards() {
	//var commission distrtypes.ValidatorAccumulatedCommission
	//commission.Commission = sdk.NewDecCoins(sdk.NewDecCoinFromDec(is.bondDenom, sdk.NewDecFromInt(sdk.NewInt(100000000000000000))))
	//currentRewards := is.network.App.DistrKeeper.GetValidatorCurrentRewards(is.ctx, is.validator.GetOperator())
	//currentRewards.Rewards = currentRewards.Rewards.Add(commission.Commission...)
	//is.network.App.DistrKeeper.SetValidatorCurrentRewards(is.ctx, is.delegation.GetValidatorAddr(), currentRewards)
	//is.network.App.DistrKeeper.SetValidatorAccumulatedCommission(is.ctx, is.delegation.GetValidatorAddr(), commission)
	//is.network.NextBlockAfter(time.Hour * 24)
	//// end block to bond validator and increase block height
	//is.network.App.StakingKeeper.BlockValidatorUpdates(is.ctx)
	//// allocate rewards to validator (of these 50% will be paid out to the delegator)
	//allocatedRewards := sdk.NewDecCoins(sdk.NewDecCoin(is.bondDenom, sdk.NewInt(20000000000000000)))
	//is.network.App.DistrKeeper.AllocateTokensToValidator(is.ctx, is.validator, allocatedRewards)
	//is.network.NextBlockAfter(time.Hour * 24)
}

func TestIntegrationSuite(t *testing.T) {
	is = new(IntegrationTestSuite)

	// Run Ginkgo integration tests
	RegisterFailHandler(Fail)
	RunSpecs(t, "Distribution Extension Suite")
}

var _ = Describe("Distribution Extension -", func() {
	var (
		distributionCallerContractAddr common.Address
		err                            error
		sender                         keyring.Key

		// contractData is a helper struct to hold the addresses and ABIs for the
		// different contract instances that are subject to testing here.
		contractData ContractData
		passCheck    testutil.LogCheckArgs

		gas      = hexutil.Uint64(900000)
		gasPrice = (*hexutil.Big)(big.NewInt(875000000))
	)

	BeforeEach(func() {
		is.SetupTest()

		// Default sender, amount
		sender = is.keyring.GetKey(0)

		distributionCallerContractAddr, err = is.factory.DeployContract(
			sender.Priv,
			evmtypes.TransactionArgs{Gas: &gas, GasPrice: gasPrice}, // NOTE: passing empty struct to use default values
			//evmtypes.TransactionArgs{}, // NOTE: passing empty struct to use default values
			factory.ContractDeploymentData{
				Contract: solc.DistributionCallerContract,
			},
		)
		Expect(err).ToNot(HaveOccurred(), "failed to deploy contract")

		contractData = ContractData{
			ownerPriv:      sender.Priv,
			precompileAddr: is.precompile.Address(),
			precompileABI:  is.precompile.Abi(),
			contractAddr:   distributionCallerContractAddr,
			contractABI:    solc.DistributionCallerContract.ABI,
		}

		passCheck = testutil.LogCheckArgs{}.WithExpPass(true)
		err = is.network.NextBlock()
		Expect(err).ToNot(HaveOccurred(), "failed to advance block")

	})

	Context("Direct precompile queries", func() {
		Context("Direct precompile queries", func() {
			It("should return the correct value", func() {
				is.setupDistributionPrecompile()
				is.AllocateRewards()
				var rewards []distribution.Coin

				txArgs, callArgs := getTxAndCallArgs(directCall, contractData, distribution.WithdrawDelegatorRewardsMethodName, is.addr, is.delegation.ValidatorAddress)
				_, ethRes, err := is.factory.CallContractAndCheckLogs(sender.Priv, txArgs, callArgs, passCheck)
				Expect(err).ToNot(HaveOccurred(), "unexpected result calling contract")

				err = is.abi.UnpackIntoInterface(&rewards, distribution.WithdrawDelegatorRewardsMethodName, ethRes.Ret)
				Expect(err).ToNot(HaveOccurred(), "failed to unpack balances")
				Expect(rewards).To(Equal(distribution.BaseRewardRes))
			})

			It("error, called by other delegator", func() {
				is.setupDistributionPrecompile()
				txArgs, callArgs := getTxAndCallArgs(directCall, contractData, distribution.WithdrawDelegatorRewardsMethodName, is.addr2, is.delegation.ValidatorAddress)
				_, _, err := is.factory.CallContractAndCheckLogs(sender.Priv, txArgs, callArgs, passCheck)
				Expect(err).To(HaveOccurred(), "error while calling the precompile")
				Expect(err.Error()).To(ContainSubstring("does not match the delegator address"), "expected different origin error")
			})
		})
	})
})
