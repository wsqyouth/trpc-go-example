package filter

import (
	"context"
	"trpc.group/trpc-go/trpc-go/codec"
	"trpc.group/trpc-go/trpc-go/filter"
	"trpc.group/trpc-go/trpc-go/log"
)

// nolint
func init() {
	filter.Register("server_filter", serverFilter, nil)
}

func serverFilter(ctx context.Context, req interface{}, next filter.ServerHandleFunc) (rsp interface{}, err error) {
	log.InfoContext(ctx, "server filter start %s", "trpc.GetMetaData(ctx, shared.AuthKey)")
	msg := codec.Message(ctx)
	md := msg.ServerMetaData()
	log.Debug("get filter msg: %v", md)

	// run business logic
	rsp, err = next(ctx, req)

	log.InfoContext(ctx, "server filter end")
	return rsp, err
}
