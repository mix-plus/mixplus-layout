package service

import (
	"context"
	"fmt"

	hello "api1/api/hello/v1"
)

func (h *HelloService) SayHello(ctx context.Context, req *hello.HelloReq) (*hello.HelloResp, error) {
	return &hello.HelloResp{
		Id:      req.Id,
		Message: fmt.Sprintf("Hello %d !", req.Id),
	}, nil
}
