package server

import (
	hello "api1/api/hello/v1"
	"api1/internal/config"
	"api1/internal/service"
	"context"
	"net/http"

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
