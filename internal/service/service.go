package service

import (
	"github.com/mix-plus/mixplus-skeleton/internal/svc"

	hello "github.com/mix-plus/mixplus-skeleton/api/hello/v1"

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
