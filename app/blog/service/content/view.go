package content

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/yira97/cnworld/app/blog/storage"
)

type ContentView struct {
	UID       string    `json:"uid"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Title     string    `json:"title"`
	Describe  string    `json:"describe"`
	BodyType  string    `json:"body_type"`
	Body      []string  `json:"body"`
	Sender    string    `json:"sender"`
	Link      []string  `json:"link"`
	Tag       []string  `json:"tag"`
}

const (
	BodyTypePlain    = "plain"
	BodyTypeMarkdown = "markdown"
)

// 在路由层调用, 执行没问题后再传给服务层
// 在路由层调用的目的是为了顺便批量检查错误
func (c *ContentView) Clean4Create() (e error) {
	c.UID = ""

	if (c.BodyType != BodyTypePlain) && (c.BodyType != BodyTypeMarkdown) {
		c.BodyType = BodyTypePlain
		e = errors.New("body type invalid")
	}

	return e
}

func (c *ContentView) Init4Create() {
	c.UID = uuid.New().String()
	now := time.Now()
	c.CreatedAt = now
	c.UpdatedAt = now
}

func (c ContentView) Model() storage.ContentModel {
	m := storage.ContentModel{
		UID:       c.UID,       // Init4Create
		CreatedAt: c.CreatedAt, // Init4Create
		UpdatedAt: c.UpdatedAt, // Init4Create
		Title:     c.Title,
		Body:      c.Body,
		BodyType:  c.BodyType,
		Sender:    c.Sender,
		Link:      c.Link,
		Tag:       c.Tag,
	}
	return m
}

func ContentView_From(m *storage.ContentModel) ContentView {
	v := ContentView{
		UID:       m.UID,
		CreatedAt: m.CreatedAt,
		UpdatedAt: m.UpdatedAt,
		Title:     m.Title,
		Body:      m.Body,
		BodyType:  m.BodyType,
		Sender:    m.Sender,
		Link:      m.Link,
		Tag:       m.Tag,
	}
	return v
}
