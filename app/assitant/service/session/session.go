package session

import (
	"encoding/json"
	"time"

	"github.com/yira97/cnworld/app/assitant/options"
	"github.com/yira97/cnworld/lib/base/crypto"
	"github.com/yira97/cnworld/lib/base/verify"
)

type UserToken struct {
	UserUID   string   `json:"user_uid"`
	UserType  string   `json:"user_type"` // super_admin common
	IssueTime string   `json:"issue_time"`
	Apps      []string `json:"apps"`
}

func NewUserToken() UserToken {
	ut := UserToken{}
	ut.IssueTime = time.Now().String()
	return ut
}

func (t UserToken) JSON() []byte {
	b, err := json.Marshal(t)
	if err != nil {
		panic(err) // DEBUG
	}
	return b
}

func (t UserToken) Encrypt() []byte {
	return t.EncryptImpl(options.GetSession().Secret)
}

func (t UserToken) EncryptImpl(secret string) []byte {
	return crypto.NewGCM_encrypt(secret, t.JSON())
}

func DecryptUserToken(cipher []byte) (*UserToken, error) {
	return DecryptUserTokenImpl(options.GetSession().Secret, cipher)
}

func DecryptUserTokenImpl(secret string, cipher []byte) (*UserToken, error) {
	plain, err := crypto.NewGCM_decrypt(secret, cipher)
	if err != nil {
		return nil, verify.ErrTokenVerifyFailed
	}

	ut := new(UserToken)
	json.Unmarshal(plain, ut)
	return ut, nil
}
