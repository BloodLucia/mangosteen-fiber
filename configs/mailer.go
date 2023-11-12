package configs

import "github.com/kelseyhightower/envconfig"

type Mailer struct {
	From     string `required:"true"`
	UserName string `required:"true"`
	PassWord string `required:"true"`
	Host     string `required:"true"`
	Port     int    `required:"true"`
}

func MailerConfig() *Mailer {
	var m Mailer
	envconfig.MustProcess("MAILER", &m)

	return &m
}
