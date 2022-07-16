// SPDX-License-Identifier: GPL-3.0
pragma solidity ^0.8.12;
pragma experimental ABIEncoderV2;

import "hardhat/console.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "./library/SafeMath.sol";
import "./library/Ownable.sol";
import "./library/Context.sol";

contract CashierDesk is Context, Ownable {
    using SafeMath for uint256;

    constructor() {}

    address private CashierTokenAddress;

    // 单位常量
    uint256 public UINT_1 = 1e18; // 单位1 ether
    uint256 public UINT_0_P_1 = 1e17; // 单位 0.1 ether
    uint256 public UINT_00_P_1 = 1e16; //单位 0.01 ether
    // 1,000,000,000,000,000,000; 等价于1 ether

    // 记录项目总收入
    uint256 private totalIncome;

    // 收银代币相关
    // 设置收银的USDT代币地址
    function setCashierTokenAddress(address cashierTokenAddress) public onlyOwner {
        CashierTokenAddress = cashierTokenAddress;
    }

    // 获取收银的USDT代币地址
    function getCashierTokenAddress() public view returns (address) {
        return CashierTokenAddress;
    }

    // 收款操作相关
    // 向用户收取usdt
    function receiveUSDT(address sender, uint256 amount) public returns(bool) {
        IERC20 token = IERC20(CashierTokenAddress);
        require(
          token.transferFrom(sender, owner(), amount),
          "receive failed"
        );
        totalIncome = totalIncome.add(amount);
        return true;
    }

    // 向用户发送usdt
    function sendUSDT(address user, uint value) private returns (bool) {
        IERC20 token = IERC20(CashierTokenAddress);
        return token.transferFrom(owner(), user, value);
    }
}