// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package di

import (
	"github.com/resource-aware-jds/worker-lib/config"
	"github.com/resource-aware-jds/worker-lib/pkg/grpc"
)

// Injectors from wire.go:

func InitializeApplication() (App, func(), error) {
	configConfig, err := config.ProvideConfig()
	if err != nil {
		return App{}, nil, err
	}
	serverConfig := config.ProvideGRPCSocketServerConfig(configConfig)
	socketServer, cleanup, err := grpc.ProvideGRPCSocketServer(serverConfig)
	if err != nil {
		return App{}, nil, err
	}
	app := ProvideApp(socketServer)
	return app, func() {
		cleanup()
	}, nil
}
