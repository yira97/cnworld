package storage

import (
	"time"

	"github.com/yira97/cnworld/lib/base/short/s_mongo"
	"github.com/yira97/cnworld/lib/base/verify"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

///////// APPLICATION

type ApplicationView struct {
	UID        string       `json:"uid"`
	Name       string       `json:"name"`
	CreatedAt  time.Time    `json:"created_at"`
	OwnerUID   string       `json:"owner_uid"`
	Permission []Permission `json:"permission"`
	Member     []Member     `json:"member"`
	Operation  []Operation  `json:"operation"`
}

type OperationDescriber struct {
	URIPrefix string   `bson:"uri_prefix" json:"uri_prefix"`
	Method    []string `bson:"method" json:"method"`
}

type Operation struct {
	Name      string               `bson:"name" json:"name"`
	Describer []OperationDescriber `bson:"describer" json:"describer"`
}

type Permission struct {
	Name      string   `bson:"name" json:"name"`           // unique
	Operation []string `bson:"operation" json:"operation"` // unique in array
}

type Member struct {
	UID        string `bson:"uid" json:"uid"`
	Permission string `bson:"permission" json:"permission"` // permission:name
}

type ApplicationStorage struct {
	// private
	MongoID    string     `json:"_id" bson:"_id,omitempty"` // omitempty，务必
	UpdatedAt  time.Time  `bson:"updated_at"`               // service generate
	DeletedAt  *time.Time `bson:"deleted_at"`               // service generate
	CreatorUID string     `bson:"creator_uid"`
	// Secret     string     `bson:"secret"` blog一类的服务, 连私网的端口就可以了.

	// protected
	UID        string       `bson:"uid"`        // service generate
	CreatedAt  time.Time    `bson:"created_at"` // service generate
	Name       string       `bson:"name"`
	OwnerUID   string       `bson:"owner_uid"`
	Permission []Permission `bson:"permission"`
	Member     []Member     `bson:"member"`
	Operation  []Operation  `bson:"operation"`
}

const (
	ApplicationStorage_UID       = verify.RawEntityUID
	ApplicationStorage_DeletedAt = verify.RawEntityDeletedAt
	ApplicationStorage_Member    = "member"
)

func (m ApplicationStorage) View() ApplicationView {
	return ApplicationView{
		UID:        m.UID,
		Name:       m.Name,
		CreatedAt:  m.CreatedAt,
		OwnerUID:   m.OwnerUID,
		Permission: m.Permission,
		Member:     m.Member,
		Operation:  m.Operation,
	}
}

func (m ApplicationStorage) Collection() *mongo.Collection {
	return GetBaseDB().Collection(ApplicationCollection)
}

func (m ApplicationStorage) DefaultOrderedFilter(fs ...bson.E) bson.D {
	return s_mongo.DefaultDeletedOrderedFilter(fs...)
}
