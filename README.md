# tgc-adapter

## 项目依赖库

- [openwallet](https://github.com/blocktree/openwallet.git)
- [eosio-adapter](https://github.com/blocktree/eosio-adapter.git)

## 如何测试

openwtester包下的测试用例已经集成了openwallet钱包体系，创建conf目录，新建TGCIO.ini文件，编辑如下内容：

```ini

#wallet api url
ServerAPI = "http://rpc.tgcchain.com:8888"
# Cache data file directory, default = "", current directory: ./data
dataDir = ""

```

## 浏览器
http://www.tgcchain.com/
