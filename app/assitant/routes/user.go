package routes

import (
	"errors"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/yira97/cnworld/app/assitant/middleware/auth"
	"github.com/yira97/cnworld/app/assitant/service/application"
	"github.com/yira97/cnworld/app/assitant/service/dashboard"
	"github.com/yira97/cnworld/app/assitant/service/session"
	"github.com/yira97/cnworld/app/assitant/service/user"
	"github.com/yira97/cnworld/app/assitant/storage"
	"github.com/yira97/cnworld/lib/base/defawt"
	"github.com/yira97/cnworld/lib/base/short/s_fiber"
	"github.com/yira97/cnworld/lib/base/verify"
)

func CreateUser(c *fiber.Ctx) error {

	v := new(storage.UserView)

	if err := c.BodyParser(v); err != nil {
		return s_fiber.CtxSetBadReq(c, err)
	}

	inserted, err := user.CreateUser(c.Context(), *v)
	if err != nil {
		if errors.Is(err, verify.ErrFieldInvalid) {
			return s_fiber.CtxSetAcc(c, defawt.RespLayout{Status: verify.RawSituationFieldInvalid, Detail: err.Error()})
		}
		return s_fiber.CtxSetIErr(c, err)
	}
	return s_fiber.CtxSetAcc(c, defawt.RespLayout{Content: *inserted})
}

func GetUser(c *fiber.Ctx) error {
	p := user.UserIdentifiableInfo{}

	userKey := c.Params(verify.RawEntityUser)
	userValue := c.Params(verify.RawEntityUser)

	if userKey != verify.RawEntityUID && userKey != verify.RawEntityEmail {
		return s_fiber.CtxSetBadReq(c, verify.ErrRouteKeyNotExist)
	}
	if userValue == "" {
		return s_fiber.CtxSetBadReq(c, verify.ErrRouteValueInvalid)
	}

	if userKey == verify.RawEntityUser {
		p.UID = &userValue
	} else if userValue == verify.RawEntityEmail {
		p.Email = &userValue
	}

	found, err := user.GetUser(c.Context(), p)
	if err != nil {
		if errors.Is(err, verify.ErrNotExist) {
			return s_fiber.CtxSetAcc(c, defawt.RespLayout{Status: verify.RawSituationNotExist, Detail: err.Error()})
		}
		return s_fiber.CtxSetIErr(c, err)
	}
	return s_fiber.CtxSetAcc(c, defawt.RespLayout{Content: *found})
}

func GetUserDashboard(c *fiber.Ctx) error {
	p := user.UserIdentifiableInfo{}

	uid := c.Params(storage.UserStorage_UID)
	if uid == "" {
		return s_fiber.CtxSetBadReq(c, fmt.Errorf("%v:%v", "uid", verify.ErrRouteValueInvalid))
	}
	p.UID = &uid

	dv, err := dashboard.GetDashBoard(c.Context(), p)
	if err != nil {
		if errors.Is(err, verify.ErrNotExist) {
			return s_fiber.CtxSetAcc(c, defawt.RespLayout{Status: verify.RawSituationNotExist})
		}
	}
	return s_fiber.CtxSetAcc(c, defawt.RespLayout{Content: *dv})
}

func UpdateUser(c *fiber.Ctx) error {
	uid := c.Params("uid")
	p := new(user.UserUpdateParams)

	if err := c.BodyParser(p); err != nil {
		return s_fiber.CtxSetBadReq(c, err)
	}

	// 只能自己改自己的
	u, ok := c.Locals(auth.LocalUser).(*session.UserToken)
	if !ok {
		return verify.ErrTokenVerifyFailed
	}
	if u.UserUID != uid {
		return s_fiber.CtxSetAcc(c, defawt.RespLayout{Status: verify.ErrUserNotExist.Error()})
	}

	v, err := user.UpdateUser(c.Context(), uid, *p)
	if err != nil {
		if errors.Is(err, verify.ErrUserNotExist) {
			return s_fiber.CtxSetAcc(c, defawt.RespLayout{Status: verify.RawSituationNotExist, Detail: err.Error()})
		}
		return s_fiber.CtxSetIErr(c, err)
	}
	return s_fiber.CtxSetAcc(c, defawt.RespLayout{Content: *v})
}

func LoginUser(c *fiber.Ctx) error {
	p := new(user.VerifyUserCreDentialParams)
	if err := c.BodyParser(p); err != nil {
		return s_fiber.CtxSetBadReq(c, err)
	}

	v, err := user.VerifyUser(c.Context(), *p)
	if err != nil {
		if errors.Is(err, verify.ErrVerifyFailed) {
			return s_fiber.CtxSetAcc(c, defawt.RespLayout{Status: verify.RawSituationFailed, Detail: err.Error()})
		}
		return s_fiber.CtxSetIErr(c, err)
	}

	token := session.NewUserToken()
	token.UserUID = v.UID
	cipher := string(token.Encrypt())

	return s_fiber.CtxSetAcc(c, defawt.RespLayout{Content: *v, Token: cipher})
}

func CreateApplication(c *fiber.Ctx) error {
	u, ok := c.Locals(auth.LocalUser).(*session.UserToken)
	if !ok {
		return verify.ErrTokenVerifyFailed
	}

	va := new(storage.ApplicationView)

	if err := c.BodyParser(va); err != nil {
		return s_fiber.CtxSetBadReq(c, err)
	}

	a, err := application.CreateApplication(c.Context(), *va, storage.UserView{UID: u.UserUID})
	if err != nil || a != nil {
		return s_fiber.CtxSetBadReq(c, err)
	}

	token := session.NewUserToken()
	token.Apps = append(token.Apps, a.UID)
	cipher := string(token.Encrypt())

	return s_fiber.CtxSetAcc(c, defawt.RespLayout{Content: *a, Token: cipher})
}
