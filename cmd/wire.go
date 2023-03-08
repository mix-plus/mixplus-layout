//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"api1/internal/config"
	"api1/internal/server"
	"api1/internal/service"
	"api1/internal/svc"

	"github.com/google/wire"
)

// initApp init app application.
func initApp(c *config.Config) (*server.AppServer, error) {
	panic(wire.Build(svc.ProviderSet, service.ProviderSet, server.ProviderSet, server.NewApp))
}
