go+vue即时通讯项目
### 前端
使用vue3 + ts构建
### 后端
gin框架
### 实现功能
1. 实时聊天，群聊
2. 接收未读消息
3. 保存/删除会话信息
4. 添加好友/加入群组
5. 创建群聊
6. 邮箱验证
7. 支持markdown格式
### 项目截图
登录
![image](https://github.com/xianren68/go_chat/assets/82361146/7dcc3fb1-f015-420d-9713-2246387de43f)
注册
![image](https://github.com/xianren68/go_chat/assets/82361146/273162ec-69cd-4aeb-8718-5ebae2ee66cb)
忘记密码/修改密码
![image](https://github.com/xianren68/go_chat/assets/82361146/8f23776d-b59a-421f-b1a3-dcedaee034f9)
好友/群组列表
![image](https://github.com/xianren68/go_chat/assets/82361146/6d54ccf9-7161-4828-ac55-338f80c136b4)
聊天/会话页面
![image](https://github.com/xianren68/go_chat/assets/82361146/d9e0d01f-3403-47d4-b3ff-247f9c4b34a2)
未读通知
![image](https://github.com/xianren68/go_chat/assets/82361146/41275279-eeba-4770-9be1-3756df77b163)
![image](https://github.com/xianren68/go_chat/assets/82361146/1e6c0332-bc9b-43e0-a079-5b500ce72473)
好友信息
![image](https://github.com/xianren68/go_chat/assets/82361146/ab2eee74-dd09-4e8e-b2fe-735e2b81f17a)
信息修改
![image](https://github.com/xianren68/go_chat/assets/82361146/7a7c2035-2e85-49c7-a573-01ca211eddec)

### 运行项目
#####  克隆项目
##### 配置配置项
在根目录下添加config.ini，添加如下配置
```
# 服务器配置
port = 9000
# 邮箱授权码
emailCode = xxxxx
# 邮箱
email = xxxxx
# mysql配置
[mysql]
user = xxxx
password = 
host = 
port = 
dbname = chat
# redis配置
[redis]
host = 
port = 6379
# 七牛云配置
[qiniu]
Zone = 2
# 七牛云密钥
AccessKey = vimdq2y7IBV1UNc7B1zf-KKEB46NCmMQZjT7nhsP
SecretKey = vlqqwGjx_1rt7kQEu3rdoSt3zMUpP7rL2GdSjx1E
Bucket = gochatimg
QiniuServer = http://s3sk3sf71.hb-bkt.clouddn.com
```
##### 下载go依赖
```
go mod tidy
```
##### 下载前端依赖（在web/user）
要有ts环境以及pnpm(网上搜一下就能搞好)
```
pnpm install
```
##### 迁移数据库
运行test目录下mai.go即可
##### 启动！
根目录
```
go run main.go
```
web/user
```
pnpm run dev
```






