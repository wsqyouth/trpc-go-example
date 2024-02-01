# trpc-go-example
对trpc-go进行实战

### 一、环境安装
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


### 二、greeter代码跑通
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

### 三、参考文档:
```azure
https://github.com/trpc-group/trpc-cmdline
https://cloud.tencent.com/developer/article/2379219
```
