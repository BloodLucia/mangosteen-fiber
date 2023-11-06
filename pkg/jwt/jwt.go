package jwt

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/kalougata/bookkeeping/pkg/config"
	"regexp"
	"time"
)

type JWT struct {
	key     []byte
	issuer  string
	expires time.Duration
}

func New(conf *config.Config) *JWT {
	return &JWT{
		key:     []byte(conf.JWT.Key),
		issuer:  conf.JWT.Issuer,
		expires: time.Duration(conf.JWT.Expires),
	}
}

type MyCustomClaims struct {
	UserId string

	jwt.RegisteredClaims
}

func (j *JWT) BuildToken(claims MyCustomClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, MyCustomClaims{
		UserId: claims.UserId,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(j.expires * time.Minute)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    j.issuer,
			Subject:   "",
			ID:        "",
			Audience:  []string{},
		},
	})

	// Sign and get the complete encoded token as a string using the key
	tokenString, err := token.SignedString(j.key)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func (j *JWT) ParseToken(tokenString string) (*MyCustomClaims, error) {
	re := regexp.MustCompile(`(?i)Bearer `)
	tokenString = re.ReplaceAllString(tokenString, "")
	token, err := jwt.ParseWithClaims(tokenString, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.key, nil
	})

	if claims, ok := token.Claims.(*MyCustomClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, err
	}
}
