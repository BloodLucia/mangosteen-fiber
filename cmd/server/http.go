package main

import (
	"fmt"
	"github.com/kalougata/bookkeeping/cmd/wire"
	"github.com/kalougata/bookkeeping/pkg/config"
	"github.com/kalougata/bookkeeping/pkg/http"
	"log"
)

func main() {
	conf := config.NewConfig()

	server, cleanup, err := wire.NewApp(conf)

	if err != nil {
		log.Panicln(err)
	}

	http.Run(server.ServerHTTP, fmt.Sprintf(":%d", conf.APP.Port))

	defer cleanup()
}
