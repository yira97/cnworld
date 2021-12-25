package storage

import (
	"time"

	"github.com/yira97/cnworld/lib/base/short/s_mongo"
	"go.mongodb.org/mongo-driver/bson"
)

type FormBlueprintStorage struct {
	// private
	MongoID   string     `json:"_id" bson:"_id,omitempty"` // omitempty，务必
	AppID     string     `json:"app_id" bson:"app_id"`
	DeletedAt *time.Time `bson:"deleted_at"` // service generate
	Creator   string     `json:"creator" bson:"creator"`
	// protected
	UID         string              `bson:"uid"`        // service generate
	CreatedAt   time.Time           `bson:"created_at"` // service generate
	Title       string              `bson:"title"`
	Description string              `bson:"description"`
	FormItems   []FormItemBluePrint `bson:"form_items"`
}

func (f FormBlueprintStorage) DefaultOrderedFilter(fs ...bson.E) bson.D {
	return s_mongo.DefaultDeletedOrderedFilter(fs...)
}

func (f FormBlueprintStorage) Collection() {
	// TODO
}

// 如果两个字段都用的是同一个
type FormItemBluePrint struct {
	Position int      `bson:"index"`
	Typ      string   `bson:"type"` // ratio, checkbox, text
	Question Question `bson:"question"`
	Answer   []string `bson:"answer"`
	Weight   int      `bosn:"weight"`
}

type Question struct {
	Title       string   `bson:"title"`
	Description string   `bson:"description"`
	Choice      []string `bson:"choice"`
}

type FormReplyItem struct {
	Index   string `bson:"index"`
	Content string `bson:"content"`
}

type FormReplyStorage struct {
	// private
	MongoID string `json:"_id" bson:"_id,omitempty"` // omitempty，务必
	// protected
	UID        string          `bson:"uid"`         // service generate
	CreatedAt  time.Time       `bson:"created_at"`  // service generate
	ReplyItems []FormReplyItem `bson:"reply_items"` //
}
