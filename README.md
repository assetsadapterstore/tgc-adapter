# tgc-adapter

## 项目依赖库

- [openwallet](https://github.com/blocktree/openwallet.git)
- [eosio-adapter](https://github.com/blocktree/eosio-adapter.git)

## 如何测试

openwtester包下的测试用例已经集成了openwallet钱包体系，创建conf目录，新建TGCIO.ini文件，编辑如下内容：

```ini

#wallet api url
ServerAPI = "http://IP:PORT"
# Cache data file directory, default = "", current directory: ./data
dataDir = ""

```
## 官方资料

### 浏览器
http://www.turingblock.io/#/dashboard

## 

创建账号接口（通过这个接口创建的账户，是可以转账的）：
curl --request GET http://XXXXX/v1/evs/uid_createaccount --data '{"new_account":"tgctest", "key_pair":"EVS5GGdCa3nUU3DF2UnJqPJuHh5LUC1yTf8nXkGpwDtgniD9sM3zj"}'
