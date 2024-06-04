package regular_test

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/zeta-chain/zetacore/testutil/contracts"
	"math/big"
	"testing"

	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/vm"
	evmtypes "github.com/evmos/ethermint/x/evm/types"
	"github.com/zeta-chain/zetacore/app"
	"github.com/zeta-chain/zetacore/precompiles/regular"
	"github.com/zeta-chain/zetacore/precompiles/regular/solc"
	"github.com/zeta-chain/zetacore/precompiles/testutil"
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
	dummyRegularCallRes1 = big.NewInt(1)

	dummyPrefix     = "zeta"
	dummyBech32Addr = "zeta1h8duy2dltz9xz0qqhm5wvcnj02upy887fyn43u"
	dummyHexAddr    = common.HexToAddress("0xB9Dbc229Bf588A613C00BEE8e662727AB8121cfE")
)

// IntegrationTestSuite is the implementation of the TestSuite interface for Regular precompile
// unit testis.
type IntegrationTestSuite struct {
	bondDenom string
	addr      common.Address
	ctx       sdk.Context

	exampleContract common.Address

	// tokenDenom is the specific token denomination used in testing the Regular precompile.
	// This denomination is used to instantiate the precompile.
	network     *network.UnitTestNetwork
	factory     factory.TxFactory
	grpcHandler grpc.Handler
	keyring     keyring.Keyring

	precompile          *regular.RegularContract
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
	sk := integrationNetwork.App.StakingKeeper
	bondDenom := sk.BondDenom(is.ctx)
	evmGenesis := *evmtypes.DefaultGenesisState()
	evmGenesis.Params.EvmDenom = "azeta"
	integrationNetwork.UpdateEvmParams(evmGenesis.Params)
	Expect(bondDenom).ToNot(BeEmpty(), "bond denom cannot be empty")

	is.bondDenom = bondDenom
	is.factory = txFactory
	is.grpcHandler = grpcHandler
	is.keyring = keyring
	is.network = integrationNetwork

	is.setupRegularPrecompile()
	is.SetDummyRegularCall()
}

func (is *IntegrationTestSuite) SetDummyRegularCall() {
	example, err := is.network.App.FungibleKeeper.DeployContract(is.ctx, contracts.ExampleMetaData)
	Expect(err).ToNot(HaveOccurred(), "failed to deploy contract")
	fmt.Println("#### deploy example contract", example)
	is.exampleContract = example
}

func (is *IntegrationTestSuite) setupRegularPrecompile() {
	pcc := regular.NewRegularContract(
		is.network.App.FungibleKeeper,
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
	RunSpecs(t, "Regular Extension Suite")
}

var _ = Describe("Regular Extension -", func() {
	var (
		regularCallerContractAddr common.Address
		err                       error
		sender                    keyring.Key

		// contractData is a helper struct to hold the addresses and ABIs for the
		// different contract instances that are subject to testing here.
		contractData ContractData
		passCheck    testutil.LogCheckArgs
	)

	BeforeEach(func() {
		is.SetupTest()

		// Default sender, amount
		sender = is.keyring.GetKey(0)

		regularCallerContractAddr, err = is.factory.DeployContract(
			sender.Priv,
			//evmtypes.TransactionArgs{Gas: &gas, GasPrice: gasPrice}, // NOTE: passing empty struct to use default values
			evmtypes.TransactionArgs{}, // NOTE: passing empty struct to use default values
			factory.ContractDeploymentData{
				Contract: solc.RegularCallerContract,
			},
		)
		Expect(err).ToNot(HaveOccurred(), "failed to deploy contract")

		contractData = ContractData{
			ownerPriv:      sender.Priv,
			precompileAddr: is.precompile.Address(),
			precompileABI:  is.precompile.Abi(),
			contractAddr:   regularCallerContractAddr,
			contractABI:    solc.RegularCallerContract.ABI,
		}

		passCheck = testutil.LogCheckArgs{}.WithExpPass(true)
		err = is.network.NextBlock()
		Expect(err).ToNot(HaveOccurred(), "failed to advance block")
	})

	Context("Stateful precompile call", func() {
		Context("Stateful precompile call", func() {
			It("should return the correct regular call result", func() {
				is.setupRegularPrecompile()
				var res *big.Int

				// Before Set 1, result should be 0
				txArgs, callArgs := getTxAndCallArgs(directCall, contractData, regular.RegularCallMethodName, "bar", is.exampleContract)
				_, ethRes, err := is.factory.CallContractAndCheckLogs(sender.Priv, txArgs, callArgs, passCheck)
				Expect(err).ToNot(HaveOccurred(), "unexpected result calling contract")
				err = is.abi.UnpackIntoInterface(&res, regular.RegularCallMethodName, ethRes.Ret)
				Expect(err).ToNot(HaveOccurred(), "failed to unpack")
				Expect(len(res.Bits())).To(Equal(0))

				// Set 1 with doSucceed, no return value for doSucceed
				txArgs, callArgs = getTxAndCallArgs(directCall, contractData, regular.RegularCallMethodName, "doSucceed", is.exampleContract)
				_, ethRes, err = is.factory.CallContractAndCheckLogs(sender.Priv, txArgs, callArgs, passCheck)
				Expect(err).ToNot(HaveOccurred(), "unexpected result calling contract")
				err = is.abi.UnpackIntoInterface(&res, regular.RegularCallMethodName, ethRes.Ret)
				Expect(err).ToNot(HaveOccurred(), "failed to unpack")
				Expect(len(res.Bits())).To(Equal(0))

				// Should be 1, after doSucceed
				txArgs, callArgs = getTxAndCallArgs(directCall, contractData, regular.RegularCallMethodName, "bar", is.exampleContract)
				_, ethRes, err = is.factory.CallContractAndCheckLogs(sender.Priv, txArgs, callArgs, passCheck)
				Expect(err).ToNot(HaveOccurred(), "unexpected result calling contract")
				err = is.abi.UnpackIntoInterface(&res, regular.RegularCallMethodName, ethRes.Ret)
				Expect(err).ToNot(HaveOccurred(), "failed to unpack")
				Expect(res).To(Equal(dummyRegularCallRes1))
			})
			It("should return the correct hex addr", func() {
				is.setupRegularPrecompile()
				var res common.Address

				txArgs, callArgs := getTxAndCallArgs(directCall, contractData, regular.Bech32ToHexAddrMethodName, dummyBech32Addr)
				_, ethRes, err := is.factory.CallContractAndCheckLogs(sender.Priv, txArgs, callArgs, passCheck)
				Expect(err).ToNot(HaveOccurred(), "unexpected result calling contract")

				err = contractData.precompileABI.UnpackIntoInterface(&res, regular.Bech32ToHexAddrMethodName, ethRes.Ret)
				Expect(err).ToNot(HaveOccurred(), "failed to unpack")
				Expect(res).To(Equal(dummyHexAddr))
			})

			It("should return the correct bech32 addr", func() {
				is.setupRegularPrecompile()
				var res string

				txArgs, callArgs := getTxAndCallArgs(directCall, contractData, regular.Bech32ifyMethodName, dummyPrefix, dummyHexAddr)
				_, ethRes, err := is.factory.CallContractAndCheckLogs(sender.Priv, txArgs, callArgs, passCheck)
				Expect(err).ToNot(HaveOccurred(), "unexpected result calling contract")

				err = contractData.precompileABI.UnpackIntoInterface(&res, regular.Bech32ifyMethodName, ethRes.Ret)
				Expect(err).ToNot(HaveOccurred(), "failed to unpack")
				Expect(res).To(Equal(dummyBech32Addr))
			})
		})
	})
})
