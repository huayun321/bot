# bot
swap bot

## 03-30
初始化项目

* 安装solc 

```shell
brew update
brew tap ethereum/ethereum
brew install solidity
```

* 安装abigen
```shell
go get -u github.com/ethereum/go-ethereum
cd $GOPATH/src/github.com/ethereum/go-ethereum/
make
make devtools
```

## 03-31
添加配置文件

## 04-04
todo
只赚一刀策略
将算出的多余利润放入GPRICE

under price提价
同nonce提价

## 04-05
20比交易测试
0.59238394
0x8c938d49ed896f03974dac63b4df205c7c39786d095976795a198794ac78e2da

todo 
heroku background app
binance smart chain node 

整理log
租服务器
整理swap配置文件

解决aws账户问题
充值bnb busd


## 04-08

优化
underprice //ok
low nonce //ok
todo
任意位数的token

watch config

full node 重新测试

竞价

日志统计

平台化 数据分析 自动配置 

gprice 自动化 get sugesst

如果不使用bnb 那么油费的计算应该要改变

写个合约


获取 最新的pair对 得到相应TOKEN
获取波动大的pair对
支持 不同小数点的 token
余额策略


优化nonce 只在初始化的时候 获取 然后放入 swap状态中


