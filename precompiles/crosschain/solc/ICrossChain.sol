pragma solidity ^0.8.7;

/// @dev The ICrossChain contract's address.
address constant ICROSSCHAIN_PRECOMPILE_ADDRESS = 0x0000000000000000000000000000000000000100;

/// @dev The ICrossChain contract's instance.
ICrossChain constant ICROSSCHAIN_CONTRACT = ICrossChain(ICROSSCHAIN_PRECOMPILE_ADDRESS);

interface ICrossChain {
  struct GasPrice {
    string creator;
    string index;
    int64 chainId;
    string[] signers;
    uint64[] blockNums;
    uint64[] prices;
    uint64 medianIndex;
  }

  function gasPrice(int64 chainID) external view returns (
    string memory creator,
    string memory index,
    int64 chainId,
    string[] memory signers,
    uint64[] memory blockNums,
    uint64[] memory prices,
    uint64 medianIndex,
    bool found
  );
}