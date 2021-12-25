package auth

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/yira97/cnworld/app/assitant/service/session"
	"github.com/yira97/cnworld/lib/base/short/s_fiber"
)

const (
	authHeaderName  = "assitant_auth"
	LocalUser       = "local_user"
	selfApplication = "assitant"
)

func New() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// 获取header
		// https://github.com/gofiber/fiber/issues/1148
		cipher := c.Get(authHeaderName)
		if cipher != "" {
			token, err := session.DecryptUserToken([]byte(cipher))
			if err != nil || token == nil {
				return s_fiber.CtxSetBadReq(c, err)
			}

			find := false
			for _, a := range token.Apps {
				if a == selfApplication {
					find = true
				}
			}
			if !find {
				return s_fiber.CtxSetIErr(c, errors.New("内部错误, applications设置不对"))
			}

			c.Locals(LocalUser, *token)
		}
		return c.Next()
	}
}
