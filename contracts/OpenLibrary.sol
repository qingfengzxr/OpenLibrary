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

    // 章节结构 -- begin
    struct Section {
        uint256 SectionNumber; //章节编号 => Book.SerialNumber + SectionCnt
        string Name; //章节名称
        string Cid;  //ipfs存储cid
        uint256 Price; // 阅读所需支付的金额
    }

    // Section.SectionNumber => (user.address => bool) 章节授权信息索引map
    mapping (uint256 => mapping(address => bool)) SectionAllowances; //授权信息
    // 章节结构 -- end

    // 书籍结构 -- begin
    struct Book {
        uint256 SerialNumber; //图书馆索引编号
        string ISBN; //国际标准书号
        string Name; //书籍名称
        string Desc; //书籍描述、简介
        address Author; //作者
        uint256 SectionCnt; //多少章节
    }

    // Book.SerialNumber => (SectionNumber => Section) 书籍章节信息索引map
    mapping(uint256 => mapping(uint256 => Section)) BookSections;  
    // 书籍结构 -- end


    // struct BookRack {
    //     address Author; //作者
    //     Book[] Books; //该作者的书籍
    // }

    // struct Library {
    //     mapping(address => BookRack) BooksTidyByAuthors; //按作者整理的书籍
    //     BookRack[] BookRacks; //所有书架
    //     mapping(string => Book) BooksTidyByName; //按书名整理的书籍
    //     Book[] Books;
    // }

    // Library private libraryInfo;

    // 组织形式，为每本书籍预留100W章
    uint256 private BookId = 1000000;
    uint256 private BookPoint = 1000000;
    mapping(uint256 => Book) private Books;


    // 作者创建书籍
    function CreateBook(string memory name, string memory desc) public returns(bool) {
        Book memory newBook;

        newBook.Author = msg.sender;
        newBook.SerialNumber = BookId;
        newBook.Name = name;
        newBook.Desc = desc;
        newBook.SectionCnt = 0; //刚创建，章节数为0

        Books[BookId] = newBook;
        BookId += BookPoint;
        return true;
    }

    // 根据书籍id，书籍对应的章节id返回书籍信息集对应章节信息
    function GetBookById(uint256 bookSerialNumber, uint256 sectionNumber) public view returns(Book memory, Section memory) {
        return (Books[bookSerialNumber], BookSections[bookSerialNumber][sectionNumber]);
    }
}