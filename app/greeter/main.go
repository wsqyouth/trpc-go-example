package main

import (
	"context"
	"fmt"
	"time"
	"trpc-go-example/proto/greeter"

	"trpc.group/trpc-go/trpc-go"
)

func main() {
	s := trpc.NewServer()
	greeter.RegisterGreeterService(s, greeterImpl{})
	_ = s.Serve()
}

type greeterImpl struct{}

func (greeterImpl) Hello(ctx context.Context, req *greeter.HelloRequest) (*greeter.HelloResponse, error) {
	fmt.Println("req:", req)
	rsp := &greeter.HelloResponse{}
	rsp.Response = fmt.Sprintf("%s to you, too", req.Greeting)
	rsp.TimestampMsec = time.Now().UnixMilli()
	return rsp, nil
}
