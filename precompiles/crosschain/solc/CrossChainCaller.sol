pragma solidity ^0.8.7;

import "./ICrossChain.sol";

contract CrossChainCaller {
// Function to get gas price by calling the getGasPrice function of the ICrossChain implementation
    function callGasPrice(int64 chainID) external view returns (
        string memory creator,
        string memory index,
        int64 chainId,
        string[] memory signers,
        uint64[] memory blockNums,
        uint64[] memory prices,
        uint64 medianIndex,
        bool found
    ) {
        // Function signature of gasPrice function
        bytes4 functionSignature = bytes4(keccak256("gasPrice(int64)"));

        // Encode chainID as int64
        bytes memory encodedChainID = abi.encodePacked(chainID);
        uint256 paddingLength = 32 - encodedChainID.length;
        bytes memory padding = new bytes(paddingLength);
        bytes memory encodedData = abi.encodePacked(functionSignature, padding, encodedChainID);

        // Call the precompiled contract at the address
        (bool success, bytes memory returnData) = address(ICROSSCHAIN_PRECOMPILE_ADDRESS).staticcall(encodedData);

        // Decode return data if successful
        if (success) {
            return abi.decode(returnData, (string, string, int64, string[], uint64[], uint64[], uint64, bool));
        } else {
            // Handle unsuccessful call (revert or exception)
            revert("Call failed");
        }
    }
}