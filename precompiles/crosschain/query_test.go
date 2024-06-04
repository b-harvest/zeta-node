package crosschain_test

import (
	"github.com/zeta-chain/zetacore/precompiles/crosschain"
)

func (p *PrecompileTestSuite) TestTotalSupply() {
	method := p.precompile.Abi().Methods[crosschain.GasPriceMethodName]

	testcases := []struct {
		name     string
		malleate func() []interface{}
		expRs    crosschain.GasPriceRes
	}{
		{
			name: "pass - exising gas prices",
			malleate: func() []interface{} {
				return []interface{}{
					existingChainId,
				}
			},
			expRs: dummyGasPriceRes,
		},
		{
			name: "fail - non-existing chain-id",
			malleate: func() []interface{} {
				return []interface{}{
					wrongChainId,
				}
			},
			expRs: emptyGasPriceRes,
		},
	}

	for _, tc := range testcases {
		tc := tc

		p.Run(tc.name, func() {
			p.SetupTest()

			bz, err := p.precompile.GasPrice(
				p.network.GetContext(),
				&method,
				tc.malleate(),
			)

			p.Require().NoError(err)
			var gasPriceRes crosschain.GasPriceRes
			err = p.precompile.Abi().UnpackIntoInterface(&gasPriceRes, method.Name, bz)
			p.Require().NoError(err)
			p.Require().Equal(tc.expRs, gasPriceRes)
		})
	}
}
