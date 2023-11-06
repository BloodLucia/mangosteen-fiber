package middleware

import (
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/gookit/goutil"
	"github.com/kalougata/bookkeeping/pkg/e"
	myJwt "github.com/kalougata/bookkeeping/pkg/jwt"
	"github.com/kalougata/bookkeeping/pkg/response"
	"log"
)

type JWTMiddleware struct {
	jwt *myJwt.JWT
}

func (jm *JWTMiddleware) JWTAuth() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		tokenString := ctx.Get("Authorization")
		if goutil.IsEqual(tokenString, "") {
			return response.Handle(ctx, e.ErrUnauthorized(), nil)
		}
		claims, err := jm.jwt.ParseToken(tokenString)
		if err != nil {
			if errors.Is(err, jwt.ErrTokenExpired) {
				return response.Handle(ctx, e.ErrUnauthorized().WithMsg("token已过期"), nil)
			}
			return response.Handle(ctx, e.ErrUnauthorized(), nil)
		}

		log.Println(fmt.Sprintf("用户ID: %s", claims.UserId))

		ctx.Set("userId", claims.UserId)

		return ctx.Next()
	}
}

func NewJWTMiddleware(jwt *myJwt.JWT) *JWTMiddleware {
	return &JWTMiddleware{jwt}
}
