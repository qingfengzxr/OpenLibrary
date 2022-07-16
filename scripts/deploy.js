async function main() {
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
}

main()
    .then(() => process.exit(0))
    .catch(error => {
        console.error(error);
        process.exit(1);
    });