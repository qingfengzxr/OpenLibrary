## 一、代币接口
代币名称：OPLBToken
燃烧及按时释放尚未实现，有待后续开发。

### 1. 代币授权相关
用户购买章节时，需要先查询是否授权了图书馆收银台合约可以操作OPLB代币。没有授权时需要让用户授权，否则无法支付。
1) 查询某接口授权数量
```
函数名：allowance
参数一：address user 用户地址
参数二：address spender 被授权使用方地址
返回值：uint256 授权数量
```

2）用户授权某个合约或者用户可以操作的代币数量
```sol
函数名：approve
参数一：address spender 被授权使用方地址
参数二：uint256 amount  授权数量
返回值：bool
```

## 二、图书馆接口
关键数据结构解释：
每本书籍有自己的编号：**SerialNumber**
这个编号的起始值为 1000000； 新增单位为：1000000；
即图书馆内新增第二本书时，其SerialNumber为：2000000
后面的100W零值数字为书籍所对应的章节预留。

每个章节有自己的**SectionNumber**
这个编号的起始值为(SerialNumber)，新增单位为：1；每本书籍单独计算起始值；
举例：
1000001 => 代表编号1000000书籍的第一章。
2000003 => 代表编号2000000书籍的第三章。

书籍信息数据结构
```C
struct Book {
    uint256 SerialNumber; //图书馆索引编号
    string ISBN; //国际标准书号
    string Name; //书籍名称
    string Desc; //书籍描述、简介
    address Author; //作者
    uint256 SectionCnt; //多少章节
    uint256 ReaderCnt; //读者数
}
```

章节信息书籍结构
```C
struct Section {
    //章节编号 => Book.SerialNumber + SectionCnt
    uint256 SectionNumber; 
    string Name; //章节名称
    string Cid;  //ipfs存储cid
    uint256 Price; // 阅读所需支付的金额
}
```

### 1.作者操作
1）创建书籍
```sol
函数名：createBook
参数一：string name 书名
参数二：string desc 描述
返回值：无
关联事件：CreateBook(address author, uint256 bookSerialNumber);
```
2）添加章节
```sol
函数名：addSection
参数一：uint256 bookSerialNumber 书籍编号
参数二：string name 章节名字
参数三：string cid 章节在ipfs上的cid
参数四：uint256 price 作者设置该章节的收取费用
返回值：无
关联事件：AddSection(address author, uint256 bookSerialNumber, uint256 sectionNumber)
```

### 2.用户操作
1）阅读具体章节
```sol
函数名：readSection
参数一：uint256 bookSerialNumber
参数二：uint256 sectionNumber
返回值：(Book, Section) 两个返回值
```
etherjs返回值示例：
```json
 [
  [
    BigNumber { value: "1000000" },
    '',
    'OPLB测试书籍',
    'OPLB测试书籍的简介，这是一个开放式的书库...',
    '0x70997970C51812dc3A010C7d01b50e0d17dc79C8',
    BigNumber { value: "1" },
    BigNumber { value: "1" },
    SerialNumber: BigNumber { value: "1000000" },
    ISBN: '',
    Name: 'OPLB测试书籍',
    Desc: 'OPLB测试书籍的简介，这是一个开放式的书库...',
    Author: '0x70997970C51812dc3A010C7d01b50e0d17dc79C8',
    SectionCnt: BigNumber { value: "1" },
    ReaderCnt: BigNumber { value: "1" }
  ],
  [
    BigNumber { value: "1000001" },
    '书籍第一章',
    'ipfs://xxxsjskljskljflkjqlkjiooq',
    BigNumber { value: "1000000000000000000" },
    SectionNumber: BigNumber { value: "1000001" },
    Name: '书籍第一章',
    Cid: 'ipfs://xxxsjskljskljflkjqlkjiooq',
    Price: BigNumber { value: "1000000000000000000" }
  ]
]
```

2) 购买书本某章节
调用用户购买书本某章节接口之前，需要查询用户是否已被授权(已经购买)
```sol
函数名：userBuySection
参数一：uint256 bookSerialNumber
参数二：uint256 sectionNumber
返回值：无
```

3）查询章节阅读权限
```sol
函数名：sectionAllowance
参数一：uint256 sectionNumber
返回值：bool
```