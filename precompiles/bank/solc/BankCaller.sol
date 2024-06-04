// SPDX-License-Identifier: LGPL-3.0-only
pragma solidity ^0.8.7;

import "./IBank.sol";

contract BankCaller {

    function callMint(address account, uint256 amount) external payable returns (bool) {
        return IBANK_CONTRACT.mint(account, amount);
    }

    function callBurn(address account, uint256 amount) external payable returns (bool) {
        return IBANK_CONTRACT.burn(account, amount);
    }

    function callBalanceOf(address account, address target) external view returns (uint256) {
        return IBANK_CONTRACT.balanceOf(account, target);
    }

    function callTransfer(address account, address target, uint256 amount) external payable returns (bool) {
        return IBANK_CONTRACT.transfer(account, target, amount);
    }
}