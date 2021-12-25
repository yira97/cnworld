package user

import (
	"context"
	"errors"
	"fmt"
	"net/mail"
	"time"
	"unicode"

	"github.com/google/uuid"
	"github.com/yira97/cnworld/app/assitant/service"
	"github.com/yira97/cnworld/app/assitant/storage"
	"github.com/yira97/cnworld/lib/base/short/s_mongo"
	"github.com/yira97/cnworld/lib/base/verify"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var _ service.CreateReady = &User{}

const (
	PasswordMinLength    = 6
	PasswordMaxLength    = 32
	NameMaxLength        = 32
	NameMinLength        = 1
	DescriptionMaxLength = 255
)

type User struct {
	V storage.UserView
	M *storage.UserStorage
}

type UserIdentifiableInfo struct {
	Email *string `json:"email"`
	UID   *string `json:"uid"`
}

func (u User) checkEmail(s string) error {
	_, e := mail.ParseAddress(s)
	if e != nil {
		return fmt.Errorf("%v:%w", e, verify.ErrEmailFieldInvalid)
	}
	return nil
}

func (u User) checkName(s string) error {
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

func (u User) checkPassword(s string) error {
	lenPw := len(s)
	if lenPw < PasswordMinLength {
		return fmt.Errorf("太短:%w", verify.ErrPasswordFieldInvalid)
	}
	if lenPw > PasswordMaxLength {
		return fmt.Errorf("太长:%w", verify.ErrPasswordFieldInvalid)
	}
	for _, c := range s {
		if !unicode.IsDigit(c) && !unicode.IsLetter(c) {
			return fmt.Errorf("存在不支持字符:%w", verify.ErrPasswordFieldInvalid)
		}
	}
	return nil
}

func (u User) checkDescription(s string) error {
	lenDesc := len(s)
	if lenDesc > DescriptionMaxLength {
		return fmt.Errorf("太长:%w", verify.ErrDescriptionFieldInvalid)
	}
	return nil
}

func (u *User) Clean4Create() {
	u.V.UID = ""
	u.V.EmailValid = false
}

func (u *User) Check4Create() (e error) {

	if e = u.checkEmail(u.V.Email); e != nil {
		return
	}

	if e = u.checkName(u.V.Name); e != nil {
		return
	}

	if e = u.checkPassword(u.V.Password); e != nil {
		return
	}

	if e = u.checkDescription(u.V.Description); e != nil {
		return
	}

	return nil
}

func (u *User) Init4Create() {
	u.V.UID = uuid.NewString()

	now := time.Now()
	u.V.CreatedAt = now
	u.V.EmailValid = false
}

// 用view创建storage
func (u *User) Create() (e error) {
	u.Clean4Create()
	if e = u.Check4Create(); e != nil {
		return
	}
	u.Init4Create()

	u.M = &storage.UserStorage{
		UID:         u.V.UID,
		CreatedAt:   u.V.CreatedAt,
		UpdatedAt:   u.V.CreatedAt, // 使用CreatedAt
		Email:       u.V.Email,
		EmailValid:  u.V.EmailValid,
		Password:    u.V.Password,
		Name:        u.V.Name,
		Gender:      u.V.Gender,
		Avatar:      u.V.Avatar,
		Description: u.V.Description,
		Birth:       u.V.Birth,
	}
	return
}

// 用参数info, 拉取最新数据, 更新view和storage
// 返回值:
//   可能错误 ErrUserNotExist
func (u *User) Sync(ctx context.Context, info UserIdentifiableInfo) error {
	filter := u.M.DefaultOrderedFilter()

	// 不光邮箱要对, 而且还是验证过的邮箱
	if info.Email != nil {
		filter = append(filter, bson.E{
			Key:   storage.UserStorage_EmailValid,
			Value: true,
		})
		filter = append(filter, bson.E{
			Key:   storage.UserStorage_Email,
			Value: *info.Email,
		})
	} else if info.UID != nil {
		filter = append(filter, bson.E{
			Key:   storage.UserStorage_UID,
			Value: *info.UID,
		})
	} else {
		return verify.ErrFieldInvalid
	}

	foundRes := u.M.Collection().FindOne(ctx, filter)
	if foundRes.Err() == mongo.ErrNoDocuments {
		return verify.ErrUserNotExist
	}
	foundRes.Decode(u.M)
	u.V = u.M.View()
	return nil
}

// 以storage的状态保存
func (u *User) Save(ctx context.Context) (e error) {
	if u.M == nil {
		return errors.New("模型为空")
	}
	_, e = u.M.Collection().InsertOne(ctx, u.M)
	if e != nil {
		return e
	}
	return e
}

func CreateUser(ctx context.Context, v storage.UserView) (*storage.UserView, error) {
	u := User{V: v}

	err := u.Create()
	if err != nil {
		return nil, err
	}

	err = u.Save(ctx)
	if err != nil {
		return nil, err
	}

	err = u.Sync(ctx, UserIdentifiableInfo{UID: &v.Email})
	if err != nil {
		return nil, err
	}
	return &u.V, nil
}

// 给外层用的, 内层请创建User对象, 然后使用Sync方法
// 返回值:
//   password字段会被清除
func GetUser(ctx context.Context, p UserIdentifiableInfo) (*storage.UserView, error) {
	u := User{}

	err := u.Sync(ctx, p)
	if err != nil {
		return nil, err
	}
	// 清除密码
	u.V.Password = ""
	return &u.V, nil
}

type VerifyUserCreDentialParams struct {
	User     UserIdentifiableInfo `json:"user"`
	Password string               `json:"password"`
}

func VerifyUser(ctx context.Context, p VerifyUserCreDentialParams) (*storage.UserView, error) {
	u := User{}
	u.Sync(ctx, p.User)
	if u.M.Password != p.Password {
		return &storage.UserView{}, verify.ErrPasswordFieldInvalid
	}
	return &u.V, nil
}

type UserUpdateParams struct {
	Name        *string `json:"name"`
	Description *string `json:"description"`
}

func UpdateUser(ctx context.Context, uid string, p UserUpdateParams) (*storage.UserView, error) {
	u := User{
		V: storage.UserView{
			UID: uid,
		},
	}
	err := u.Sync(ctx, UserIdentifiableInfo{UID: &u.M.UID})
	if err != nil {
		return nil, err
	}

	change := bson.M{}

	if p.Name != nil {
		if u.checkName(*p.Name) == nil {
			change[storage.UserStorage_Name] = *p.Name
		}
	}

	if p.Description != nil {
		if u.checkDescription(*p.Description) == nil {
			change[storage.UserStorage_Description] = *p.Description
		}
	}

	update := bson.M{
		"$set": change,
	}

	opt := s_mongo.OptionFindOneAndUpdateThenReturn()

	result := u.M.Collection().FindOneAndUpdate(ctx, u.M.DefaultOrderedFilter(), update, &opt)

	if result.Err() != nil {
		return nil, result.Err()
	}

	// Decode the result
	decodeErr := result.Decode(&u.M)
	if decodeErr != nil {
		return nil, decodeErr
	}

	u.V = u.M.View()

	return &u.V, nil
}
