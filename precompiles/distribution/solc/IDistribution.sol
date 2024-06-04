// SPDX-License-Identifier: LGPL-3.0-only
pragma solidity ^0.8.7;

import "./CosmosTypes.sol";

/// @dev The DistributionI contract's address.
address constant DISTRIBUTION_PRECOMPILE_ADDRESS = 0x0000000000000000000000000000000000000200;

/// @dev Define all the available distribution methods.
string constant MSG_WITHDRAW_DELEGATOR_REWARD = "/cosmos.distribution.v1beta1.MsgWithdrawDelegatorReward";

/// @dev The DistributionI contract's instance.
DistributionI constant DISTRIBUTION_CONTRACT = DistributionI(
  DISTRIBUTION_PRECOMPILE_ADDRESS
);

/// @author Evmos
/// @title Distribution Precompile Contract PoC
/// @dev The interface through which solidity contracts will interact with Distribution
/// @custom:address 0x0000000000000000000000000000000000000801
interface DistributionI {
  // ...

  /// @dev ClaimRewards defines an Event emitted when rewards are claimed
  /// @param delegatorAddress the address of the delegator
  /// @param amount the amount being claimed
  event ClaimRewards(
    address indexed delegatorAddress,
    uint256 amount
  );

  /// @dev WithdrawDelegatorRewards defines an Event emitted when rewards from a delegation are withdrawn
  /// @param delegatorAddress the address of the delegator
  /// @param validatorAddress the address of the validator
  /// @param amount the amount being withdrawn from the delegation
  event WithdrawDelegatorRewards(
    address indexed delegatorAddress,
    address indexed validatorAddress,
    uint256 amount
  );

  /// TRANSACTIONS

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
