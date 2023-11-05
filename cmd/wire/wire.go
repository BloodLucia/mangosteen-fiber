//go:build wireinject
// +build wireinject

package wire

import (
	"github.com/google/wire"
	"github.com/kalougata/bookkeeping/internal/controller"
	"github.com/kalougata/bookkeeping/internal/data"
	"github.com/kalougata/bookkeeping/internal/server"
)

func NewApp() (*server.Server, func(), error) {
	panic(wire.Build(
		data.NewData,
		controller.NewAuthController,
		server.NewHTTPServer,
		server.NewServer,
	))
}
