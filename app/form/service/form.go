package service

import (
	"time"

	"github.com/yira97/cnworld/app/form/storage"
	"github.com/yira97/cnworld/lib/base/short/s_mongo"
)

type Form struct {
	UID         string     `json:"uid"`
	CreatedAt   time.Time  `bson:"created_at"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	FormItems   []FormItem `json:"form_items"`
}

type FormItem struct {
	BpIndex  int              `json:"bp_index"`
	Typ      string           `bson:"type"` // ratio, checkbox, text
	Question storage.Question `bson:"question"`
	Weight   int              `bosn:"weight"`
}

func GenerateForm(blueprintUID string) {
	filter := s_mongo.DefaultDeletedOrderedFilter()

	filter = filter
	// TODO
}
