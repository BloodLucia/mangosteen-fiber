package configs

import "github.com/kelseyhightower/envconfig"

type App struct {
	Name string `required:"true"`
	Port int    `required:"true"`
}

func AppConfig() *App {
	var app App
	envconfig.MustProcess("APP", &app)

	return &app
}
