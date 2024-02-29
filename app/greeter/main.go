package main

import (
	"context"
	"fmt"
	"time"
	_ "trpc-go-example/app/greeter/filter"
	"trpc-go-example/proto/greeter"
	"trpc.group/trpc-go/trpc-go"
	"trpc.group/trpc-go/trpc-go/codec"
	"trpc.group/trpc-go/trpc-go/log"
	// 加载插件包
	"trpc-go-example/app/greeter/config"
)

func main() {
	s := trpc.NewServer()
	greeter.RegisterGreeterService(s, greeterImpl{})
	_ = s.Serve()
}

type greeterImpl struct{}

func (greeterImpl) Hello(ctx context.Context, req *greeter.HelloRequest) (*greeter.HelloResponse, error) {
	// 可以看到控制台 warn 级别以上的日志都会被打印出来, 日志中是info级别
	log.Tracef("trace msg:%s", req)
	log.Debugf("debug msg:%s", req)
	log.Infof("info msg:%s", req)
	log.Warnf("warn msg:%s", req)
	log.Errorf("error msg:%s", req)

	rsp := &greeter.HelloResponse{}
	rsp.Response = fmt.Sprintf("%s to you, too", req.Greeting)
	rsp.TimestampMsec = time.Now().UnixMilli()

	config.Record() // 通过包访问公共函数
	log.Debugf("SayHi config: %+v", config.DefaultAppConfig)

	// 测试全局变量
	globalConf := trpc.GlobalConfig()
	log.Warnf("globalConf: %+v", globalConf)

	// 测试获取metadata
	msg := codec.Message(ctx)
	md := msg.ServerMetaData()
	for k, v := range md {
		log.Debugf("get metadata key : %s, value : %s", k, string(v))
	}
	return rsp, nil
}
