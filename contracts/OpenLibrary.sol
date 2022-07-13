// SPDX-License-Identifier: GPL-3.0
pragma solidity ^0.8.12;
pragma experimental ABIEncoderV2;

import "hardhat/console.sol";
import "./IBEP20.sol";
import "./library/SafeMath.sol";
import "./library/Ownable.sol";
import "./library/Context.sol";

contract OpenLibrary is Context, Ownable {
    using SafeMath for uint256;

    struct Section {
        uint256 SectionNumber; //章节编号
        string Name; //章节名称
        string Cid;  //ipfs存储cid
        uint256 Price; // 阅读所需支付的金额
        mapping (address => bool) Allowances; //授权信息
    }

    struct Book {
        string SerialNumber; //图书馆索引编号
        string ISBN; //国际标准书号
        string Name; //书籍名称
        string Desc; //书籍描述、简介
        address Author; //作者
        Section[] Sections; //章节信息
        uint256 SectionCnt; //多少章节
    }

    struct BookRack {
        address Author; //作者
        Book[] Books; //该作者的书籍
    }

    struct Library {
        mapping(address => BookRack) BooksTidyByAuthors; //按作者整理的书籍
        BookRack[] BookRacks; //所有书架
        mapping(string => Book) BooksTidyByName; //按书名整理的书籍
    }
}