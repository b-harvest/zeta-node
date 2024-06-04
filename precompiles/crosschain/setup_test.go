package crosschain_test

import (
	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/zeta-chain/zetacore/app"
	"github.com/zeta-chain/zetacore/precompiles/crosschain"
	"testing"

	"github.com/stretchr/testify/suite"
	"github.com/zeta-chain/zetacore/testutil/integration/zeta/factory"
	"github.com/zeta-chain/zetacore/testutil/integration/zeta/grpc"
	testkeyring "github.com/zeta-chain/zetacore/testutil/integration/zeta/keyring"
	"github.com/zeta-chain/zetacore/testutil/integration/zeta/network"
)

var s *PrecompileTestSuite

// PrecompileTestSuite is the implementation of the TestSuite interface for ERC20 precompile
// unit tests.
type PrecompileTestSuite struct {
	suite.Suite

	bondDenom, tokenDenom string

	// tokenDenom is the specific token denomination used in testing the ERC20 precompile.
	// This denomination is used to instantiate the precompile.
	network     *network.UnitTestNetwork
	factory     factory.TxFactory
	grpcHandler grpc.Handler
	keyring     testkeyring.Keyring

	precompile          *crosschain.CrossChainContract
	precompiledContract vm.PrecompiledContract
	abi                 abi.ABI

	kvGasConfig storetypes.GasConfig
	cdc         codec.Codec
}

func TestPrecompileTestSuite(t *testing.T) {
	s = new(PrecompileTestSuite)
	suite.Run(t, s)
}

func (p *PrecompileTestSuite) SetupTest() {
	encodingConfig := app.MakeEncodingConfig()
	keyring := testkeyring.New(2)
	integrationNetwork := network.NewUnitTestNetwork(
		network.WithPreFundedAccounts(keyring.GetAllAccAddrs()...),
	)
	grpcHandler := grpc.NewIntegrationHandler(integrationNetwork)
	txFactory := factory.New(integrationNetwork, grpcHandler)

	ctx := integrationNetwork.GetContext()
	sk := integrationNetwork.App.StakingKeeper
	bondDenom := sk.BondDenom(ctx)
	p.Require().NotEmpty(bondDenom, "bond denom cannot be empty")

	p.kvGasConfig = storetypes.TransientGasConfig()
	p.cdc = encodingConfig.Codec

	p.bondDenom = bondDenom
	p.factory = txFactory
	p.grpcHandler = grpcHandler
	p.keyring = keyring
	p.network = integrationNetwork

	p.setupCrossChainPrecompile()
	p.SetDummyGasPrice()
}

func (p *PrecompileTestSuite) SetDummyGasPrice() {
	p.network.App.CrosschainKeeper.SetGasPrice(p.network.GetContext(), dummyGasPrice)
}

func (p *PrecompileTestSuite) setupCrossChainPrecompile() {
	pcc := crosschain.NewCrossChainContract(
		p.network.App.CrosschainKeeper,
		p.cdc,
		p.kvGasConfig,
	)
	p.precompile = pcc
	p.abi = p.precompile.Abi()
}
