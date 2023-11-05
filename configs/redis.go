package configs

import "github.com/kelseyhightower/envconfig"

type Redis struct {
	Host   string `default:"127.0.0.1"`
	Port   int    `default:"6379"`
	User   string `required:"true"`
	Passwd string `default:""`
	Db     int    `required:"true"`
}

func RedisConfig() *Redis {
	var rdb Redis
	envconfig.MustProcess("RDB", &rdb)

	return &rdb
}
