package configs

import "github.com/kelseyhightower/envconfig"

type JWT struct {
	Key    string `required:"true"`
	Issuer string `required:"true"`
}

func JWTConfig() *JWT {
	var jwt JWT
	envconfig.MustProcess("JWT", &jwt)

	return &jwt
}
