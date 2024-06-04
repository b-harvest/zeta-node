package crosschain_test

import (
	"math/big"
	"testing"

	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/zeta-chain/zetacore/app"
	"github.com/zeta-chain/zetacore/precompiles/crosschain"
	"github.com/zeta-chain/zetacore/precompiles/crosschain/solc"
	"github.com/zeta-chain/zetacore/precompiles/testutil"
	"github.com/zeta-chain/zetacore/x/crosschain/types"

	"github.com/ethereum/go-ethereum/common"
	evmtypes "github.com/evmos/ethermint/x/evm/types"
	"github.com/zeta-chain/zetacore/testutil/integration/zeta/factory"
	"github.com/zeta-chain/zetacore/testutil/integration/zeta/grpc"
	"github.com/zeta-chain/zetacore/testutil/integration/zeta/network"

	"github.com/zeta-chain/zetacore/testutil/integration/zeta/keyring"

	//nolint:revive // dot imports are fine for Ginkgo
	. "github.com/onsi/ginkgo/v2"
	//nolint:revive // dot imports are fine for Ginkgo
	. "github.com/onsi/gomega"
)

var is *IntegrationTestSuite
var (
	existingChainId = int64(1)
	wrongChainId    = int64(10000)
	dummyCreator    = "zeta1rx9r8hff0adaqhr5tuadkzj4e7ns2ntg446vtt"
	dummyGasPrice   = types.GasPrice{
		Creator:     dummyCreator,
		Index:       "",
		ChainId:     existingChainId,
		Signers:     []string{dummyCreator},
		BlockNums:   []uint64{1, 2},
		Prices:      []uint64{1, 2},
		MedianIndex: 1,
	}

	dummyGasPriceRes = crosschain.GasPriceRes{
		Creator:     dummyGasPrice.Creator,
		Index:       dummyGasPrice.Index,
		ChainId:     dummyGasPrice.ChainId,
		Signers:     dummyGasPrice.Signers,
		BlockNums:   dummyGasPrice.BlockNums,
		Prices:      dummyGasPrice.Prices,
		MedianIndex: dummyGasPrice.MedianIndex,
		Found:       true,
	}

	emptyGasPriceRes = crosschain.GasPriceRes{
		//..
		Signers:   []string{},
		BlockNums: []uint64{},
		Prices:    []uint64{},
		Found:     false,
	}
)

// IntegrationTestSuite is the implementation of the TestSuite interface for CrossChain precompile
// unit testis.
type IntegrationTestSuite struct {
	bondDenom string
	addr      common.Address

	// tokenDenom is the specific token denomination used in testing the CrossChain precompile.
	// This denomination is used to instantiate the precompile.
	network     *network.UnitTestNetwork
	factory     factory.TxFactory
	grpcHandler grpc.Handler
	keyring     keyring.Keyring

	precompile          *crosschain.CrossChainContract
	precompiledContract vm.PrecompiledContract
	abi                 abi.ABI

	kvGasConfig storetypes.GasConfig
	cdc         codec.Codec
}

const (
	CallGasPrice = "callGasPrice"
)

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

	ctx := integrationNetwork.GetContext()
	sk := integrationNetwork.App.StakingKeeper
	bondDenom := sk.BondDenom(ctx)
	evmGenesis := *evmtypes.DefaultGenesisState()
	evmGenesis.Params.EvmDenom = "azeta"
	integrationNetwork.UpdateEvmParams(evmGenesis.Params)
	Expect(bondDenom).ToNot(BeEmpty(), "bond denom cannot be empty")

	is.bondDenom = bondDenom
	is.factory = txFactory
	is.grpcHandler = grpcHandler
	is.keyring = keyring
	is.network = integrationNetwork

	is.setupCrossChainPrecompile()
	is.SetDummyGasPrice()
}

func (is *IntegrationTestSuite) SetDummyGasPrice() {
	is.network.App.CrosschainKeeper.SetGasPrice(is.network.GetContext(), dummyGasPrice)
	val, found := is.network.App.CrosschainKeeper.GetGasPrice(is.network.GetContext(), dummyGasPrice.ChainId)
	Expect(found).ToNot(BeFalse())
	Expect(val).To(Equal(dummyGasPrice))
}

func (is *IntegrationTestSuite) setupCrossChainPrecompile() {
	pcc := crosschain.NewCrossChainContract(
		is.network.App.CrosschainKeeper,
		is.cdc,
		is.kvGasConfig,
	)

	is.precompile = pcc
	is.abi = is.precompile.Abi()
}

func TestIntegrationSuite(t *testing.T) {
	is = new(IntegrationTestSuite)

	// Run Ginkgo integration tests
	RegisterFailHandler(Fail)
	RunSpecs(t, "CrossChain Extension Suite")
}

var _ = Describe("CrossChain Extension -", func() {
	var (
		crosschainCallerContractAddr common.Address
		err                          error
		sender                       keyring.Key

		// contractData is a helper struct to hold the addresses and ABIs for the
		// different contract instances that are subject to testing here.
		contractData ContractData
		passCheck    testutil.LogCheckArgs

		gas      = hexutil.Uint64(900000)
		gasPrice = (*hexutil.Big)(big.NewInt(875000000))
		//gas      = hexutil.Uint64(1000000)
		//gasPrice = (*hexutil.Big)(big.NewInt(875000000))
	)

	BeforeEach(func() {
		is.SetupTest()

		// Default sender, amount
		sender = is.keyring.GetKey(0)

		crosschainCallerContractAddr, err = is.factory.DeployContract(
			sender.Priv,
			evmtypes.TransactionArgs{Gas: &gas, GasPrice: gasPrice}, // NOTE: passing empty struct to use default values
			//evmtypes.TransactionArgs{}, // NOTE: passing empty struct to use default values
			factory.ContractDeploymentData{
				Contract: solc.CrossChainCallerContract,
			},
		)
		Expect(err).ToNot(HaveOccurred(), "failed to deploy contract")

		contractData = ContractData{
			ownerPriv:      sender.Priv,
			precompileAddr: is.precompile.Address(),
			precompileABI:  is.precompile.Abi(),
			contractAddr:   crosschainCallerContractAddr,
			contractABI:    solc.CrossChainCallerContract.ABI,
		}

		passCheck = testutil.LogCheckArgs{}.WithExpPass(true)
		err = is.network.NextBlock()
		Expect(err).ToNot(HaveOccurred(), "failed to advance block")
	})

	Context("Direct precompile queries", func() {
		Context("Direct precompile queries", func() {
			It("should return the correct gas price", func() {
				is.setupCrossChainPrecompile()
				var gasPriceRes crosschain.GasPriceRes

				txArgs, callArgs := getTxAndCallArgs(directCall, contractData, crosschain.GasPriceMethodName, dummyGasPrice.ChainId)
				_, ethRes, err := is.factory.CallContractAndCheckLogs(sender.Priv, txArgs, callArgs, passCheck)
				Expect(err).ToNot(HaveOccurred(), "unexpected result calling contract")

				err = is.abi.UnpackIntoInterface(&gasPriceRes, crosschain.GasPriceMethodName, ethRes.Ret)
				Expect(err).ToNot(HaveOccurred(), "failed to unpack balances")
				Expect(gasPriceRes).To(Equal(dummyGasPriceRes))
			})

			It("querying not existing chain id, should return not found with empty result", func() {
				txArgs, callArgs := getTxAndCallArgs(directCall, contractData, crosschain.GasPriceMethodName, wrongChainId)
				_, ethRes, err := is.factory.CallContractAndCheckLogs(sender.Priv, txArgs, callArgs, passCheck)
				Expect(err).ToNot(HaveOccurred(), "unexpected result calling contract")

				var gasPriceRes crosschain.GasPriceRes
				err = is.abi.UnpackIntoInterface(&gasPriceRes, crosschain.GasPriceMethodName, ethRes.Ret)
				Expect(err).ToNot(HaveOccurred(), "failed to unpack gas price")
				Expect(gasPriceRes).To(Equal(emptyGasPriceRes))
			})
		})
	})

	Context("Calls from a contract", func() {
		Context("Calls from a contract", func() {
			It("should return the correct gas price 3 tx failed. VmError: execution reverted,", func() {
				queryArgs, balancesArgs := getTxAndCallArgs(contractCall, contractData, CallGasPrice, dummyGasPrice.ChainId)
				queryArgs.Gas = &gas
				_, ethRes, err := is.factory.CallContractAndCheckLogs(sender.Priv, queryArgs, balancesArgs, passCheck)
				Expect(err).ToNot(HaveOccurred(), "unexpected result calling contract")

				var gasPriceRes crosschain.GasPriceRes
				err = is.abi.UnpackIntoInterface(&gasPriceRes, crosschain.GasPriceMethodName, ethRes.Ret)
				Expect(err).ToNot(HaveOccurred(), "failed to unpack balances")
			})
		})
	})
})
