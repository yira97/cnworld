package user

import (
	"errors"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/yira97/cnworld/app/blog/storage"
)

type TempUserView struct {
	UID       string    `json:"uid"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	UserUID   string    `json:"user_uid"`
}

func (v *TempUserView) Clean4Create() (e error) {
	v.UID = ""
	v.UserUID = ""
	if !strings.Contains(v.Email, "@") {
		return errors.New("email invalid")
	}
	return
}

func (v *TempUserView) Init4Create() {
	v.UID = uuid.New().String()
	now := time.Now()
	v.CreatedAt = now
	v.UpdatedAt = now
}

func (v TempUserView) Model() storage.UserContactModel {
	m := storage.UserContactModel{
		CreatedAt: v.CreatedAt,
		UpdatedAt: v.UpdatedAt,
		UID:       v.UID,
		Name:      v.Name,
		Email:     v.Email,
		UserUID:   v.UserUID,
	}
	return m
}

func TempUserView_From(m *storage.UserContactModel) TempUserView {
	v := TempUserView{
		UID:       m.UID,
		Name:      m.Name,
		Email:     m.Email,
		UserUID:   m.UserUID,
		CreatedAt: m.CreatedAt,
		UpdatedAt: m.UpdatedAt,
	}
	return v
}

func (v TempUserView) IsLinked() bool {
	return v.UserUID == ""
}
