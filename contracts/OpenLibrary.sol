// SPDX-License-Identifier: GPL-3.0
pragma solidity ^0.8.12;
pragma experimental ABIEncoderV2;

import "hardhat/console.sol";
import "./IBEP20.sol";
import "./library/SafeMath.sol";
import "./library/Ownable.sol";
import "./library/Context.sol";
import "./CashierDesk.sol";

contract OpenLibrary is Context, Ownable {
    using SafeMath for uint256;
    CashierDesk private cashierDesk;
    
    event CreateBook(address indexed author, uint256 indexed bookSerialNumber);
    event AddSection(address indexed author, uint256 indexed bookSerialNumber, uint256 indexed sectionNumber);


    // 设置收银台合约
    function setCashierDesk(address cashier) public onlyOwner {
        cashierDesk = CashierDesk(cashier);
    }

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
        uint256 ReaderCnt; //读者数
    }

    // Book.SerialNumber => (SectionNumber => Section) 书籍章节信息索引map
    mapping(uint256 => mapping(uint256 => Section)) BookSections;
    // 书籍结构 -- end
    //记录作者有多少本书，address => Book.SerialNumber
    mapping(address => uint256[]) authorBooks;

    // 组织形式，为每本书籍预留100W章
    uint256 private BookId = 1000000;
    uint256 private BookPoint = 1000000;
    // BookId => Book 等价于 Book.SerialNumber => Book
    mapping(uint256 => Book) private Books;


    // 作者创建书籍
    function createBook(string memory name, string memory desc) public {
        Book memory newBook;

        newBook.Author = msg.sender;
        newBook.SerialNumber = BookId;
        newBook.Name = name;
        newBook.Desc = desc;
        newBook.SectionCnt = 0; //刚创建，章节数为0

        Books[BookId] = newBook;
        BookId += BookPoint;

        // 记录这个作者新增了一本书
        authorBooks[msg.sender].push(newBook.SerialNumber);
        
        emit CreateBook(msg.sender, newBook.SerialNumber);
    }

    // 作者添加章节
    function addSection(uint256 bookSerialNumber, string memory name, string memory cid, uint256 price) public {
        Section memory newSection;
        
        newSection.Name = name;
        newSection.Cid = cid;
        newSection.Price = price;

        uint256 bookSectionCnt = Books[bookSerialNumber].SectionCnt;
        // 添加新的章节，章节数目加1
        newSection.SectionNumber = bookSerialNumber + bookSectionCnt + 1;
        // 将书籍与新增章节关联
        BookSections[bookSerialNumber][newSection.SectionNumber] = newSection;

        // 关联后书籍章节数量正式+1
        Books[bookSerialNumber].SectionCnt += 1;
        
        emit AddSection(msg.sender, bookSerialNumber, newSection.SectionNumber);
    }

    // 阅读
    // 根据书籍id，书籍对应的章节id返回书籍信息集对应章节信息
    function readSection(uint256 bookSerialNumber, uint256 sectionNumber) public view returns(Book memory, Section memory) {
        require(
            SectionAllowances[sectionNumber][msg.sender] == true,
            "sorry, you are not buy this section"
        );
        
        return (Books[bookSerialNumber], BookSections[bookSerialNumber][sectionNumber]);
    }

    // 用户购买某书章节
    function userBuySection(uint256 bookSerialNumber, uint256 sectionNumber) public {
        // 收款作者指定的金额
        require (
            cashierDesk.receiveUSDT(_msgSender(), BookSections[bookSerialNumber][sectionNumber].Price),
            "pay section fee failed."
        );

        // 购买书籍其实就是授权
        SectionAllowances[sectionNumber][msg.sender] = true;
        // 授权成功后， 读者数量+1
        Books[bookSerialNumber].ReaderCnt += 1;
    }

    // 查询用户是否有某个章节的权限
    function sectionAllowance(uint256 sectionNumber) public view returns(bool) {
        return SectionAllowances[sectionNumber][msg.sender];
    }
}