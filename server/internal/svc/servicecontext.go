package svc

import (
	"iot/server/internal/config"
	"iot/server/internal/device"
)

type ServiceContext struct {
	Config  config.Config
	Manager *device.Manager
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:  c,
		Manager: device.NewManager(),
	}
}
