package bech32_test

import (
	"fmt"
	"github.com/zeta-chain/zetacore/precompiles/bech32"
	cmn "github.com/zeta-chain/zetacore/precompiles/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/zeta-chain/zetacore/app"
)

func (s *PrecompileTestSuite) TestHexToBech32() {
	// setup basic test suite
	s.SetupTest()

	method := s.precompile.Methods[bech32.HexToBech32Method]

	testCases := []struct {
		name        string
		malleate    func() []interface{}
		postCheck   func(data []byte)
		expError    bool
		errContains string
	}{
		{
			"fail - invalid args length",
			func() []interface{} {
				return []interface{}{}
			},
			func(data []byte) {},
			true,
			fmt.Sprintf(cmn.ErrInvalidNumberOfArgs, 2, 0),
		},
		{
			"fail - invalid hex address",
			func() []interface{} {
				return []interface{}{
					"",
					"",
				}
			},
			func(data []byte) {},
			true,
			"invalid hex address",
		},
		{
			"fail - invalid bech32 HRP",
			func() []interface{} {
				return []interface{}{
					s.keyring.GetAddr(0),
					"",
				}
			},
			func(data []byte) {},
			true,
			"invalid bech32 human readable prefix (HRP)",
		},
		{
			"pass - valid hex address and valid bech32 HRP",
			func() []interface{} {
				return []interface{}{
					s.keyring.GetAddr(0),
					app.Bech32PrefixAccAddr,
				}
			},
			func(data []byte) {
				args, err := s.precompile.Unpack(bech32.HexToBech32Method, data)
				s.Require().NoError(err, "failed to unpack output")
				s.Require().Len(args, 1)
				addr, ok := args[0].(string)
				s.Require().True(ok)
				s.Require().Equal(s.keyring.GetAccAddr(0).String(), addr)
			},
			false,
			"",
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			s.SetupTest()

			bz, err := s.precompile.HexToBech32(&method, tc.malleate())

			if tc.expError {
				s.Require().Error(err)
				s.Require().ErrorContains(err, tc.errContains, err.Error())
				s.Require().Empty(bz)
			} else {
				s.Require().NoError(err)
				s.Require().NotEmpty(bz)
				tc.postCheck(bz)
			}
		})
	}
}

func (s *PrecompileTestSuite) TestBech32ToHex() {
	// setup basic test suite
	s.SetupTest()

	method := s.precompile.Methods[bech32.Bech32ToHexMethod]

	testCases := []struct {
		name        string
		malleate    func() []interface{}
		postCheck   func(data []byte)
		expError    bool
		errContains string
	}{
		{
			"fail - invalid args length",
			func() []interface{} {
				return []interface{}{}
			},
			func(data []byte) {},
			true,
			fmt.Sprintf(cmn.ErrInvalidNumberOfArgs, 1, 0),
		},
		{
			"fail - empty bech32 address",
			func() []interface{} {
				return []interface{}{
					"",
				}
			},
			func(data []byte) {},
			true,
			"invalid bech32 address",
		},
		{
			"fail - invalid bech32 address",
			func() []interface{} {
				return []interface{}{
					app.Bech32PrefixAccAddr,
				}
			},
			func(data []byte) {},
			true,
			fmt.Sprintf("invalid bech32 address: %s", app.Bech32PrefixAccAddr),
		},
		{
			"fail - decoding bech32 failed",
			func() []interface{} {
				return []interface{}{
					app.Bech32PrefixAccAddr + "1",
				}
			},
			func(data []byte) {},
			true,
			"decoding bech32 failed",
		},
		{
			"fail - invalid address format",
			func() []interface{} {
				return []interface{}{
					sdk.AccAddress(make([]byte, 256)).String(),
				}
			},
			func(data []byte) {},
			true,
			"address max length is 255",
		},
		{
			"success - valid bech32 address",
			func() []interface{} {
				return []interface{}{
					s.keyring.GetAccAddr(0).String(),
				}
			},
			func(data []byte) {
				args, err := s.precompile.Unpack(bech32.Bech32ToHexMethod, data)
				s.Require().NoError(err, "failed to unpack output")
				s.Require().Len(args, 1)
				addr, ok := args[0].(common.Address)
				s.Require().True(ok)
				s.Require().Equal(s.keyring.GetAddr(0), addr)
			},
			false,
			"",
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			s.SetupTest()

			bz, err := s.precompile.Bech32ToHex(&method, tc.malleate())

			if tc.expError {
				s.Require().Error(err)
				s.Require().ErrorContains(err, tc.errContains)
				s.Require().Empty(bz)
			} else {
				s.Require().NoError(err)
				s.Require().NotEmpty(bz)
				tc.postCheck(bz)
			}
		})
	}
}
