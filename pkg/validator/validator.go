package validator

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gookit/validate"
	"github.com/kalougata/bookkeeping/pkg/e"
)

func Checker(ctx *fiber.Ctx, data any) error {
	if err := ctx.BodyParser(data); err != nil {
		return e.ErrFormatJSON()
	}
	if v := validate.Struct(data); !v.Validate() {
		return e.ErrBadRequest().WithMsg(v.Errors.Error())
	}

	return nil
}
