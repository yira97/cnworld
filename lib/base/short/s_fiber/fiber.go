package s_fiber

import (
	"github.com/gofiber/fiber/v2"
	"github.com/yira97/cnworld/lib/base/defawt"
	"github.com/yira97/cnworld/lib/base/format"
	"github.com/yira97/cnworld/lib/base/verify"
)

func QueryInt(c *fiber.Ctx, key string, defaultValue int) int {
	return format.ToInt(c.Query(key), defaultValue)
}

func CtxSetBadReq(c *fiber.Ctx, err error) error {
	return c.Status(fiber.StatusBadRequest).SendString(err.Error())
}

func CtxSetIErr(c *fiber.Ctx, err error) error {
	return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
}

func CtxSetAcc(c *fiber.Ctx, data defawt.RespLayout) error {
	if data.Status == "" {
		data.Status = verify.RawSituationOk
	}
	return c.Status(fiber.StatusCreated).JSON(data)
}
