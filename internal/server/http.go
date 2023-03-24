package server

import (
	"context"
	"net/http"

	hello "github.com/mix-plus/mixplus-skeleton/api/hello/v1"
	"github.com/mix-plus/mixplus-skeleton/internal/config"
	"github.com/mix-plus/mixplus-skeleton/internal/service"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
)

func NewHttpServer(c *config.Config, srv *service.HelloService) *http.Server {
	hs := &http.Server{}
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	// Register gRPC server endpoint
	// Note: Make sure the gRPC server is running properly and accessible
	mux := runtime.NewServeMux()
	err := hello.RegisterHelloHandlerServer(ctx, mux, srv)
	if err != nil {
		panic(err)
	}
	hs.Addr = c.ApiConf.Addr
	hs.Handler = mux

	return hs
}
