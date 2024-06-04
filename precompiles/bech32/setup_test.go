package bech32_test

import (
	"github.com/zeta-chain/zetacore/precompiles/bech32"
	"testing"

	"github.com/stretchr/testify/suite"
	testkeyring "github.com/zeta-chain/zetacore/testutil/integration/zeta/keyring"
	"github.com/zeta-chain/zetacore/testutil/integration/zeta/network"
)

var s *PrecompileTestSuite

// PrecompileTestSuite is the implementation of the TestSuite interface for ERC20 precompile
// unit tests.
type PrecompileTestSuite struct {
	suite.Suite

	network *network.UnitTestNetwork
	keyring testkeyring.Keyring

	precompile *bech32.Precompile
}

func TestPrecompileTestSuite(t *testing.T) {
	s = new(PrecompileTestSuite)
	suite.Run(t, s)
}

func (s *PrecompileTestSuite) SetupTest() {
	keyring := testkeyring.New(2)
	integrationNetwork := network.NewUnitTestNetwork(
		network.WithPreFundedAccounts(keyring.GetAllAccAddrs()...),
	)

	s.keyring = keyring
	s.network = integrationNetwork

	precompile, err := bech32.NewPrecompile(6000)
	s.Require().NoError(err, "failed to create bech32 precompile")

	s.precompile = precompile
}
