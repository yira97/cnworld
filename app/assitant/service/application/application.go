package application

import (
	"context"
	"errors"
	"fmt"
	"time"
	"unicode"

	"github.com/google/uuid"
	"github.com/yira97/cnworld/app/assitant/service"
	"github.com/yira97/cnworld/app/assitant/storage"
	"github.com/yira97/cnworld/lib/base/verify"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var _ service.CreateReady = &Application{}

const (
	NameMaxLength = 32
	NameMinLength = 1
)

type Application struct {
	V storage.ApplicationView
	M *storage.ApplicationStorage
}

type ApplicationIdentifiableInfo struct {
	UID *string `json:"uid"`
}

type ApplicationFilterInfo struct {
	Creator *string `json:"creator"`
	Member  *string `json:"member"`
}

func (a Application) checkName(s string) error {
	lenName := len(s)
	if lenName < NameMinLength {
		return fmt.Errorf("太短:%w", verify.ErrNameFieldInvalid)
	}
	if lenName > NameMaxLength {
		return fmt.Errorf("太长:%w", verify.ErrNameFieldInvalid)
	}
	if unicode.IsDigit([]rune(s)[0]) {
		return fmt.Errorf("首位不能数字:%w", verify.ErrNameFieldInvalid)
	}
	for _, c := range s {
		if !unicode.IsDigit(c) && !unicode.IsLetter(c) && c != '_' {
			return fmt.Errorf("存在不支持字符:%w", verify.ErrNameFieldInvalid)
		}
	}
	return nil
}

func (u *Application) Clean4Create() {
	u.V.UID = ""
}

func (u *Application) Check4Create() (e error) {

	if e = u.checkName(u.V.Name); e != nil {
		return
	}

	return nil
}

func (u *Application) Init4Create() {
	u.V.UID = uuid.NewString()

	now := time.Now()
	u.V.CreatedAt = now
}

// 只要uid
func (a *Application) Create(vu storage.UserView) (e error) {
	a.Clean4Create()
	if e = a.Check4Create(); e != nil {
		return
	}
	a.Init4Create()

	a.M = &storage.ApplicationStorage{
		UpdatedAt:  a.V.CreatedAt, // 使用CreatedAt
		CreatorUID: vu.UID,
		UID:        a.V.UID,
		CreatedAt:  a.V.CreatedAt,
		Name:       a.V.Name,
		OwnerUID:   vu.UID,
		Permission: []storage.Permission{},
		Member: []storage.Member{
			{UID: vu.UID},
		},
		Operation: []storage.Operation{},
	}
	return
}

// 用info提供的信息, 从数据库拉取数据, 复制
func (a *Application) Sync(ctx context.Context, info ApplicationIdentifiableInfo) error {
	filter := a.M.DefaultOrderedFilter()

	if info.UID != nil {
		filter = append(filter, bson.E{
			Key:   storage.ApplicationStorage_UID,
			Value: info.UID,
		})
	} else {
		return verify.ErrFieldInvalid
	}

	foundRes := a.M.Collection().FindOne(ctx, filter)
	if foundRes.Err() == mongo.ErrNoDocuments {
		return verify.ErrApplicationNotExist
	}
	foundRes.Decode(a.M)
	a.V = a.M.View()
	return nil
}

// 以storage的状态保存
func (a *Application) Save(ctx context.Context) (e error) {
	if a.M == nil {
		return errors.New("模型为空")
	}
	_, e = a.M.Collection().InsertOne(ctx, a.M)
	if e != nil {
		return e
	}
	return e
}

// vu只要uid
func CreateApplication(ctx context.Context, va storage.ApplicationView, vu storage.UserView) (*storage.ApplicationView, error) {
	a := Application{V: va}

	err := a.Create(vu)
	if err != nil {
		return nil, err
	}

	err = a.Save(ctx)
	if err != nil {
		return nil, err
	}

	err = a.Sync(ctx, ApplicationIdentifiableInfo{UID: &a.V.UID})
	if err != nil {
		return nil, err
	}
	return &a.V, nil
}

func GetApplication(ctx context.Context, p ApplicationIdentifiableInfo) (*storage.ApplicationView, error) {
	a := Application{}

	err := a.Sync(ctx, ApplicationIdentifiableInfo{UID: p.UID})
	if err != nil {
		return nil, err
	}
	return &a.V, nil
}
