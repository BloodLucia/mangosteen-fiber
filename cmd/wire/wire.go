//go:build wireinject
// +build wireinject

package wire

import (
	"github.com/google/wire"
	"github.com/kalougata/bookkeeping/internal/server"
)

func NewApp() (*server.Server, func(), error) {
	panic(wire.Build(
		server.NewHTTPServer,
		server.NewServer,
	))
}
