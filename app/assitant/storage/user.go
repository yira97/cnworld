package storage

import (
	"time"

	"github.com/yira97/cnworld/lib/base/short/s_mongo"
	"github.com/yira97/cnworld/lib/base/verify"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

//////// USER

type UserView struct {
	UID         string    `json:"uid"`
	CreatedAt   time.Time `json:"created_at"`
	Email       string    `json:"email"`
	EmailValid  bool      `json:"email_valid"`
	Password    string    `json:"password"`
	Name        string    `json:"name"`
	Gender      string    `json:"gender"`
	Avatar      string    `json:"avatar"`
	Birth       string    `json:"birth"`
	Description string    `json:"description"`
}

// birth:
//   1997 year-only
//   1997-07 year-month
//   1997-07-28 year-month-day
type UserStorage struct {
	// prvate
	MongoID   string     `json:"_id" bson:"_id,omitempty"` // omitempty，务必
	UpdatedAt time.Time  `bson:"updated_at"`               // service generate
	DeletedAt *time.Time `bson:"deleted_at"`               // service generate
	Password  string     `bson:"password"`
	// protected
	UID         string    `bson:"uid"`        // service generate
	CreatedAt   time.Time `bson:"created_at"` // service generate
	Email       string    `bson:"email"`
	EmailValid  bool      `bson:"email_valid"`
	Name        string    `bson:"name"`
	Gender      string    `bson:"gender"`
	Avatar      string    `bson:"avatar"`
	Birth       string    `bson:"birth"`
	Description string    `bson:"description"`
}

const (
	UserStorage_DeletedAt   = verify.RawEntityDeletedAt
	UserStorage_MongoID     = verify.RawEntityMongoID
	UserStorage_UID         = verify.RawEntityUID
	UserStorage_Name        = verify.RawEntityName
	UserStorage_Description = verify.RawEntityDescription
	UserStorage_Email       = verify.RawEntityEmail
	UserStorage_EmailValid  = "email_valid"
)

func (m UserStorage) Collection() *mongo.Collection {
	return GetBaseDB().Collection(UserCollection)
}

func (m UserStorage) DefaultOrderedFilter() bson.D {
	return s_mongo.DefaultDeletedOrderedFilter()
}

func (m UserStorage) View() UserView {
	return UserView{
		UID:         m.UID,
		CreatedAt:   m.CreatedAt,
		Email:       m.Email,
		EmailValid:  m.EmailValid,
		Password:    m.Password,
		Name:        m.Name,
		Gender:      m.Gender,
		Avatar:      m.Avatar,
		Birth:       m.Birth,
		Description: m.Description,
	}
}
