package service

import (
	"api1/internal/svc"

	hello "api1/api/hello/v1"

	"github.com/google/wire"
)

// ProviderSet is server providers.
var ProviderSet = wire.NewSet(NewHelloServer)

type HelloService struct {
	hello.UnimplementedHelloServer

	svcCtx *svc.ServiceContext
}

func NewHelloServer(ctx *svc.ServiceContext) *HelloService {
	return &HelloService{
		svcCtx: ctx,
	}
}
