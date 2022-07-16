const { inputToConfig } = require("@ethereum-waffle/compiler");
const { expect } = require("chai");
const { ethers } = require("hardhat");


describe("BlackList contract", function () {
    let owner;
    let author;
    let user1;
    let user2;

    let openLb;
    let OPLBToken;
    let oplbToken;

    beforeEach(async function () {
        // 获取测试用户
        [owner, author, user1, user2] = await ethers.getSigners();

        // step 0: deploy cashierDesk
        CashierDesk = await ethers.getContractFactory("CashierDesk");
        cashierDesk = await CashierDesk.deploy();
        console.log("[beforeEach]CashierDesk contract address:", cashierDesk.address);

        // step 1: deploy OpenLibrary
        OpenLibrary = await ethers.getContractFactory("OpenLibrary");
        openLb = await OpenLibrary.deploy();
        console.log("[beforeEach]OpenLibrary contract address:", openLb.address);

        // step2: deploy token
        OPLBToken = await ethers.getContractFactory("OPLBToken");
        oplbToken = await OPLBToken.deploy();
        console.log("[beforeEach]OPLBToken contract address:", oplbToken.address);

        // 设置收银代币
        await cashierDesk.setCashierTokenAddress(oplbToken.address);
        // open library 设置收银台合约的地址
        await openLb.setCashierDesk(cashierDesk.address);

        usdtTokenAddress = await cashierDesk.getCashierTokenAddress();
        console.log("[beforeEach]ownerAddress: ", owner.address);
        console.log("[beforeEach]user1Address: ", user1.address);
        console.log("[beforeEach]user2Address: ", user2.address);
    });


    describe("作者相关功能测试", function () {
        it("创建书籍", async function () {
            let openLBSingerAuthor = await openLb.connect(author);
            let openLBSingerUser1 = await openLb.connect(user1);
            let openLBSingerUser2 = await openLb.connect(user2);

            // 创建第一本书
            let bookName = "OPLB测试书籍";
            let bookDesc = "OPLB测试书籍的简介，这是一个开放式的书库...";
            let blockInfo = await openLBSingerAuthor.createBook(bookName, bookDesc);
            let ret = await openLb.queryFilter("CreateBook", blockInfo.blockHash);

            expect(ret[0].args.author).to.equal(author.address);
            expect(ret[0].args.bookSerialNumber).to.equal(1000000);
            console.log("book1 test done");

            // 创建第二本书
            bookName = "OPLB测试书籍2";
            bookDesc = "OPLB测试书籍2的简介，这是一个开放式的书库...";
            blockInfo = await openLBSingerAuthor.createBook(bookName, bookDesc);
            ret = await openLb.queryFilter("CreateBook", blockInfo.blockHash);
            expect(ret[0].args.author).to.equal(author.address);
            expect(ret[0].args.bookSerialNumber).to.equal(2000000);
            console.log("book2 test done");

            // 用户1也创建一本书
            bookName = "OPLB测试书籍3";
            bookDesc = "OPLB测试书籍3的简介，这是一个开放式的书库...";
            blockInfo = await openLBSingerUser1.createBook(bookName, bookDesc);
            ret = await openLb.queryFilter("CreateBook", blockInfo.blockHash);
            expect(ret[0].args.author).to.equal(user1.address);
            expect(ret[0].args.bookSerialNumber).to.equal(3000000);
            console.log("book3 test done");
        });
        
        it("添加章节", async function () {
            let openLBSingerAuthor = await openLb.connect(author);
            let openLBSingerUser1 = await openLb.connect(user1);
            let openLBSingerUser2 = await openLb.connect(user2);

            // 创建第一本书
            let bookName = "OPLB测试书籍";
            let bookDesc = "OPLB测试书籍的简介，这是一个开放式的书库...";
            let blockInfo = await openLBSingerAuthor.createBook(bookName, bookDesc);
            let ret = await openLb.queryFilter("CreateBook", blockInfo.blockHash);

            expect(ret[0].args.author).to.equal(author.address);
            expect(ret[0].args.bookSerialNumber).to.equal(1000000);
            console.log("book1 test done");

            // 添加一个章节
            let sectionName = "书籍第一章";
            let sectionCid = "ipfs://xxxsjskljskljflkjqlkjiooq";
            // 阅读需要收取一个代币
            let sectionPrice = ethers.utils.parseEther("1.0");
            blockInfo = await openLBSingerAuthor.addSection(ret[0].args.bookSerialNumber, sectionName, sectionCid, sectionPrice);
            ret = await openLb.queryFilter("AddSection", blockInfo.blockHash);

            expect(ret[0].args.author).to.equal(author.address);
            expect(ret[0].args.bookSerialNumber).to.equal(1000000);
            expect(ret[0].args.sectionNumber).to.equal(1000001);
            console.log("add section test done");
        })
    });
    
    describe("用户相关功能测试", function () {
        it("查看书籍", async function () {
            let openLBSingerAuthor = await openLb.connect(author);
            let openLBSingerUser1 = await openLb.connect(user1);
            let openLBSingerUser2 = await openLb.connect(user2);

            // 创建第一本书
            let bookName = "OPLB测试书籍";
            let bookDesc = "OPLB测试书籍的简介，这是一个开放式的书库...";
            let blockInfo = await openLBSingerAuthor.createBook(bookName, bookDesc);
            let ret = await openLb.queryFilter("CreateBook", blockInfo.blockHash);

            expect(ret[0].args.author).to.equal(author.address);
            expect(ret[0].args.bookSerialNumber).to.equal(1000000);
            console.log("book1 test done");

            // 添加一个章节
            let sectionName = "书籍第一章";
            let sectionCid = "ipfs://xxxsjskljskljflkjqlkjiooq";
            // 阅读需要收取一个代币
            let sectionPrice = ethers.utils.parseEther("1.0");
            blockInfo = await openLBSingerAuthor.addSection(ret[0].args.bookSerialNumber, sectionName, sectionCid, sectionPrice);
            ret = await openLb.queryFilter("AddSection", blockInfo.blockHash);

            expect(ret[0].args.author).to.equal(author.address);
            expect(ret[0].args.bookSerialNumber).to.equal(1000000);
            expect(ret[0].args.sectionNumber).to.equal(1000001);
            console.log("add section test done");

            // 给用户转点钱
            let tokenAsSignerOwner = await oplbToken.connect(owner);
            let tokenAsSignerUser1 = await oplbToken.connect(user1);
            let tokenAsSignerUser2 = await oplbToken.connect(user2);

            // 授权owner可以操作10000个代币
            await oplbToken.approve(owner.address, ethers.utils.parseEther("10000.0"));
            // 授权user1可以操作10000个代币
            // await oplbToken.approve(user1.address, ethers.utils.parseEther("10000.0"));
            // user1 授权收银台合约可以操作10000个代币
            tokenAsSignerUser1.approve(cashierDesk.address, ethers.utils.parseEther("10000.0"));

            // 向user1转2000个代币
            const tranferToUser1Amount = ethers.utils.parseEther("2000.0");
            await tokenAsSignerOwner.transferFrom(owner.address, user1.address, tranferToUser1Amount);
            let user1Balance = await tokenAsSignerUser1.balanceOf(user1.address);
            console.log("user1Balance: ", user1Balance);
            // 查看书籍之前用户需要购买书籍
            await openLBSingerUser1.userBuySection(1000000, 1000001);

            // 用户查看书籍
            let bookInfo = await openLBSingerUser1.getBookBySerialNumber(1000000, 1000001);
            console.log("bookInfo: ", bookInfo);
        });
    })
});