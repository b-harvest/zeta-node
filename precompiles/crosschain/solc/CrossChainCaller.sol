pragma solidity ^0.8.7;

import "./ICrossChain.sol";

contract CrossChainCaller {
    // Function to get gas price by calling the getGasPrice function of the ICrossChain implementation
    function callGasPrice(int64 chainID) public view returns (
        string memory creator,
        string memory index,
        int64 chainId,
        string[] memory signers,
        uint64[] memory blockNums,
        uint64[] memory prices,
        uint64 medianIndex,
        bool found
    ) {
        return ICROSSCHAIN_CONTRACT.gasPrice(chainID);
    }
}