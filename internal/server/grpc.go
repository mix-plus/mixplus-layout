package server

import (
	"api1/internal/config"
	"api1/internal/service"
	"api1/internal/svc"

	hello "api1/api/hello/v1"

	"github.com/mix-plus/go-mixplus/mrpc"
	"google.golang.org/grpc"
)

func NewGrpcServer(c *config.Config, svc *svc.ServiceContext) *mrpc.RpcServer {

	srv := service.NewHelloServer(svc)
	s := mrpc.MustNewServer(c.RpcServerConf, func(g *grpc.Server) {
		// grpc register
		hello.RegisterHelloServer(g, srv)
	})

	return s
}
