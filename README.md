# trpc-go-example
>对trpc-go进行实战
参考文章: https://cloud.tencent.com/developer/article/2384591
--------------
### 第一部分
#### 一、环境安装
需要安装trpc-cmdline项目里的readme相关文件，主要包括:
```azure
Install protoc
Install flatc
Install protoc-gen-go
Install goimports
Install mockgen
Install protoc-gen-validate and protoc-gen-validate-go
之后安装:
go install trpc.group/trpc-go/trpc-cmdline/trpc@latest
trpc version
```


#### 二、greeter代码跑通
定义pb
make pb生成桩文件
修改业务代码,定义yaml
运行服务端代码:
````
// 注意是在main.go文件路径下
go run main.go -conf trpc_go.yaml
````
运行客户端请求,支持Get和Post:
```azure
curl -i "http://127.0.0.1:8000/trpc.greeter.Greeter/Hello?greeting=Morning"
    
curl http://127.0.0.1:8000/trpc.greeter.Greeter/Hello --header 'Content-Type:application/json' -d '{"greeting":"Good afternoon"}'
```

#### 三、参考文档:
```azure
https://github.com/trpc-group/trpc-cmdline
https://cloud.tencent.com/developer/article/2379219
```
--------------

### 第二部分
#### 2.1 遇到的问题
当rpc文件开始引入其他的proto文件时,编译会报错,我先使用手动的办法生成
```azure
分别进入rpc/user rpc/httpauth目录下命令生成并处理
protoc -I . -I ../ --go_out=. *.proto
```
参考文档: https://www.jianshu.com/p/a0286b58098e

解决办法：这样是不行的, 正确解决办法是protodir指定import的proto文件

#### 2.2 直接参考andrew的github工程先跑起来整个项目
主要分为两部分: http前端服务, rpc后端服务
项目架构其实设计的不怎么好, 后续再优化，第一阶段是运行项目成功。
步骤:
```azure
1. 定义proto文件, make pb生成
2. 写user server服务代码,然后将yaml中定义为http protocol，用curl测试：
    curl -i    http://127.0.0.1:8002/demo.account.User/GetAccountByUserName    -H 'Content-Type: application/json'   -H 'cache-control: no-cache'    -H 'X-Client-Id: 1111655025'   -d '{"username":"benben4652"}'
3. 将server服务的yaml 中定义为trpc protocol服务启动
    cd trpc-go-example/app/user
    go run .  -conf conf/trpc_go.yaml
4. 启动http cgi服务,注意检查yaml配置文件
    cd trpc-go-example/app/http-auth-server
    go run .  -conf conf/trpc_go.yaml
5. 串联测试
    curl '127.0.0.1:8001/demo/auth/Login?username=coopers'
```

--------------
### 第三部分 添加插件
给第一个greeter项目添加一个配置插件, 一般项目启动后一些账号密码是放到配置文件中的,我们需要在项目启动时加载
配置文件信息到全局变量，之后在调用中使用。
这里参考: https://github.com/trpc-group/trpc-go/blob/main/examples/features/plugin/custom_plugin.go 实现
步骤：
```azure
1. 定义config及自定义插件 [在greeter项目添加conf] 
2. 在main.go中引入包
3. 在yaml中添加自定义配置
    
验证:
go run main.go -conf trpc_go.yaml      
curl -i "http://127.0.0.1:8000/trpc.greeter.Greeter/Hello?greeting=Morning"
```

同理,添加日志插件
参考: https://github.com/trpc-group/trpc-go/tree/main/examples/features/log
```azure
1. 在yaml中添加日志配置
2. 日志包引用了即可。这个内部的init已封装了zap日志实现了，不需要我们自己开发
```

同理,测试了下yaml文件中的global全局变量,这个可以用来获取服务的全局信息: 测试正式环境。

--------------
### 第四部分 添加filter
参考: https://github.com/trpc-group/trpc-go/tree/main/examples/features/filter
```azure
代码见greeter项目
```