### lotus client doc
#### lotus client import
- cmd line `lotus client import /root/file`
- rpc request
```
{
    "id": 1,
    "jsonrpc": "2.0",
    "method": "Filecoin.ClientImport",
    "params": [{
        "Path": "/root/file",
        "IsCAR": false,
    }]
}
```
- rpc response
```
{
    "id": 1,
    "jsonrpc": "2.0",
    "result": {
        "Root": {
             "/": "bafykbzaceb43hehfjrw5he7oifyhejttrh36bpchs6yn3cwe5ip5swdgwhhzy",
         },
        "ImportID": 1657851778366660827,
    }
    "error": {}
}
```

#### lotus client deal
- cmd line
```
lotus client deal --verfied-deal --from=t3rkwb5rc4fgeqohfvcfvxlevj3s5uh4gjddbgjmquurcc3wmbr6c273cns54joylyyrthpsjkoixu72jwrx5a bafykbzaceb43hehfjrw5he7oifyhejttrh36bpchs6yn3cwe5ip5swdgwhhzy t038342 0.005 51840
--from参数指定用于支付的钱包地址
其中第一个参数为数据cid，第二个参数为接受交易的矿工ID，第三个参数为交易费用，第四个参数为交易约定时长，最少180天。
```
- rpc request
```
{
    "id": 1,
    "jsonrpc": "2.0",
    "method": "Filecoin.ClientStartDeal",
    "params": [{
        "Data": {
            "TransferType": "graphsync",
            "Root": "bafykbzaceb43hehfjrw5he7oifyhejttrh36bpchs6yn3cwe5ip5swdgwhhzy",
        },
        "Wallet": "t3rkwb5rc4fgeqohfvcfvxlevj3s5uh4gjddbgjmquurcc3wmbr6c273cns54joylyyrthpsjkoixu72jwrx5a",
        "Miner": "t038342",
        "EpochPrice": 0.005,
        "MinBlockDuration": 51840,
        "DealStartEpoch": -1,
        "FastRetrieval": true,
        "VerifiedDeal": false,
        "ProviderCollateral": 0
    }]
}
```
- rpc response
```
{
    "id": 1,
    "jsonrpc": "2.0",
    "result": {"/": "bafykbzacebtqkg6ycah3x3ph47vsepnpgbi5cmwm2r2wdzxar6a5sqfk7jvds"},
    "error": {}
}
```
```
DealStartEpoch 指定为-1表示从当前时间开始
VerifiedDeal 可以默认指定为false，表示未经过验证的数据（未经过验证的数据就没有对矿工的10倍算力的奖励）
响应数据为交易的cid
响应的错误信息结构为
{
    "code": number,
    "message": string,
    "data": object
}
``` 