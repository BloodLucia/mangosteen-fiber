package response

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/kalougata/bookkeeping/pkg/e"
	"net/http"
)

const (
	SUCCESS = 1
	FAIL    = -1
)

type response struct {
	Code    int    `json:"code"`
	Success bool   `json:"success"`
	Msg     string `json:"msg"`
	Data    any    `json:"data"`
}

func Handle(c *fiber.Ctx, err error, data any) error {
	if err == nil {
		return c.JSON(&response{
			Code:    SUCCESS,
			Success: true,
			Msg:     "success",
			Data:    data,
		})
	}

	var myErr *e.Error

	if !errors.As(err, &myErr) {
		c.SendStatus(http.StatusInternalServerError)
		return c.JSON(&response{
			Code:    FAIL,
			Success: false,
			Msg:     "Unknown Error",
			Data:    nil,
		})
	}

	c.SendStatus(myErr.Code)
	return c.JSON(&response{
		Code:    FAIL,
		Success: false,
		Msg:     myErr.Msg,
		Data:    data,
	})
}
