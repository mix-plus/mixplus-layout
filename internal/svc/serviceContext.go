package svc

import (
	"api1/internal/config"

	"github.com/google/wire"
)

// ProviderSet is server providers.
var ProviderSet = wire.NewSet(NewServiceContext)

type ServiceContext struct {
	Config *config.Config
}

func NewServiceContext(c *config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
	}
}
