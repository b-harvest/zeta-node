// SPDX-License-Identifier: LGPL-3.0-only
pragma solidity ^0.8.7;

import "./CosmosTypes.sol" as types;
import "./IDistribution.sol" as distribution;

contract DistributionCaller {
    // ...

    function testWithdrawDelegatorRewardsFromContract(
        string memory _valAddr
    ) public returns (types.Coin[] memory) {
        return
        distribution.DISTRIBUTION_CONTRACT.withdrawDelegatorRewards(
            address(this),
            _valAddr
        );
    }

    function testWithdrawDelegatorRewards(
        address _delAddr,
        string memory _valAddr
    ) public returns (types.Coin[] memory) {
        return
        distribution.DISTRIBUTION_CONTRACT.withdrawDelegatorRewards(
            _delAddr,
            _valAddr
        );
    }

    // ...
}
