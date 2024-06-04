// SPDX-License-Identifier: LGPL-3.0-only
pragma solidity ^0.8.7;

import "./CosmosTypes.sol";

/// @dev The DistributionI contract's address.
address constant DISTRIBUTION_PRECOMPILE_ADDRESS = 0x0000000000000000000000000000000000000068; // 104

/// @dev The DistributionI contract's instance.
DistributionI constant DISTRIBUTION_CONTRACT = DistributionI(
  DISTRIBUTION_PRECOMPILE_ADDRESS
);

/// @author Evmos
/// @title Distribution Precompile Contract PoC
/// @dev The interface through which solidity contracts will interact with Distribution
interface DistributionI {
  // ...

  /// @dev Withdraw the rewards of a delegator from a validator
  /// @param delegatorAddress The address of the delegator
  /// @param validatorAddress The address of the validator
  /// @return amount The amount of Coin withdrawn
  function withdrawDelegatorRewards(
    address delegatorAddress,
    string memory validatorAddress
  ) external returns (Coin[] calldata amount);

  // ...
}
