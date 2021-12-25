package verify

import (
	"errors"
	"fmt"
)

const (
	///////////////////////////////////////// entity/behave/situation
	Entity               = "entity"
	RawEntityUser        = "user"
	RawEntityApplication = "application"
	RawEntityName        = "name"
	RawEntityEmail       = "email"
	RawEnityPassword     = "password"
	RawEntityDescription = "description"
	RawEntityUID         = "uid"
	RawEntityDeletedAt   = "deleted_at"
	RawEntityMongoID     = "_id"
	RawEntityToken       = "token"
	RawEntityValue       = "value"
	RawEntityRouteValue  = "route_value"
	RawEntityRouteKey    = "route_key"
	// behave
	RawBehaveVerify = "verify"
	// situation
	RawSituationOk           = "ok"
	RawSituationNotExist     = "not-exist"
	RawSituationFieldInvalid = "field-invalid"
	RawSituationFailed       = "failed"
)

var (
	////////////////////////////////////////// error(situation)
	//
	ErrNotExist            = errors.New(RawSituationNotExist)
	ErrUserNotExist        = fmt.Errorf("%s:%w", RawEntityUser, ErrNotExist)
	ErrApplicationNotExist = fmt.Errorf("%s:%w", RawEntityApplication, ErrNotExist)
	ErrRouteKeyNotExist    = fmt.Errorf("%s:%w", RawEntityRouteKey, ErrNotExist)
	//
	ErrFieldInvalid            = errors.New(RawSituationFieldInvalid)
	ErrNameFieldInvalid        = fmt.Errorf("%s:%w", RawEntityName, ErrFieldInvalid)
	ErrEmailFieldInvalid       = fmt.Errorf("%s:%w", RawEntityEmail, ErrFieldInvalid)
	ErrPasswordFieldInvalid    = fmt.Errorf("%s:%w", RawEnityPassword, ErrFieldInvalid)
	ErrDescriptionFieldInvalid = fmt.Errorf("%s:%w", RawEntityDescription, ErrFieldInvalid)
	ErrRouteValueInvalid       = fmt.Errorf("%s:%w", RawEntityRouteValue, ErrFieldInvalid)
	//
	ErrFailed            = errors.New(RawSituationFailed)
	ErrVerifyFailed      = fmt.Errorf("%s:%w", RawBehaveVerify, ErrFailed)
	ErrTokenVerifyFailed = fmt.Errorf("%s:%w", RawEntityToken, ErrVerifyFailed)
	ErrUserVerifyFailed  = fmt.Errorf("%s:%w", RawEntityUser, ErrVerifyFailed)
)
