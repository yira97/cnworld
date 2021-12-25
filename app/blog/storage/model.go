package storage

import (
	"time"
)

///////////////////////////////////

const (
	CommonColumn__ID         = "_id"
	CommonColumn_UID         = "uid"
	CommonColumn_CreatedAt   = "created_at"
	CommonColumn_UpdatedAt   = "updated_at"
	CommonColumn_DeletedAt   = "deleted_at"
	CommonColumn_Name        = "name"
	CommonColumn_Email       = "email"
	CommonColumn_Content     = "content"
	CommonColumn_Title       = "title"
	CommonColumn_Sender      = "sender"
	CommonColumn_Receiver    = "receiver"
	CommonColumn_Description = "description"
	CommonColumn_MessageUID  = "message_uid"
	CommonColumn_HasRead     = "has_read"
	CommonColumn_MarkedAs    = "marked_as"
	CommonColumn_ContentUID  = "content_uid"
)

type UserContactModel struct {
	MongoID   string    `json:"_id" bson:"_id,omitempty"` // mongodb generate
	CreatedAt time.Time `bson:"created_at"`               // service generate
	UpdatedAt time.Time `bson:"updated_at"`               // service generate
	// 优先级 user_uid > email > name
	// 所有的contact 必须要有一个 name
	// email是对其他用户隐藏的
	Name  string `bson:"name"`
	Email string `bson:"email"`
	// 对于已经绑定了user_uid的UserContactInfo, 强制要求用户以user身份登录, 系统以User的信息来解释信息, 如name相关的属性
	UserUID string `bson:"user_uid"`
}

//  Message是类似邮件一样的中低频率信息,而不是短信一样的高频信息, 这里的结构一旦创建, 不会再被改变.
//  收件人和发件人都必须是人, 而不是组织.
type MessageModel struct {
	ID        string    `json:"_id" bson:"_id,omitempty"` // mongodb generate
	CreatedAt time.Time `bson:"created_at"`               // service generate
	UID       string    `bson:"uid"`                      // service generate
	Title     string    `bson:"title"`
	Content   string    `bson:"content"`
	// 发件人和收件人不能是同一个人
	Sender   string `bson:"sender"`
	Receiver string `bson:"receiver"`
	// 为了减轻服务器负担, 系统的邮件都带保质期, 同时这个属性也对用户提供一种功能, 实效信息.
	AccessibleBefore time.Time `bson:"accessible_before"`
}

// 因收件人的操作而改变.
type MetaMessageModel struct {
	ID         string     `json:"_id" bson:"_id,omitempty"` // mongodb generate
	UpdatedAt  time.Time  `bson:"updated_at"`               // service generate
	DeletedAt  *time.Time `bson:"deleted_at"`               // service generate
	MessageUID string     `bson:"message_uid"`
	HasRead    bool       `bson:"has_read"`
	MarkedAs   string     `bson:"marked_as"`
}

type ContentModel struct {
	ID        string    `json:"_id" bson:"_id,omitempty"` // mongodb generate
	CreatedAt time.Time `bson:"created_at"`               // service generate
	UpdatedAt time.Time `bson:"updated_at"`               // service generate
	// https://stackoverflow.com/questions/43653402/how-can-i-assign-a-null-value-to-date-field-using-mongo-go-driver-instead-of-dat
	// 用这个形式指代可以为空
	DeletedAt *time.Time `bson:"deleted_at"` // service generate
	UID       string     `bson:"uid"`        // service generate
	Title     string     `bson:"title"`
	// 为了避免含义混淆, 这里不叫content, 叫body
	// 数组的原因是, 有些时候content分成好几个段落, 与前端渲染相关, 这里前端展示优先.
	// 保证body的存储结构一致后, 前端自己根据返回的body_type来判断应该如何渲染.
	Body []string `bson:"body"`
	// 'PLAIN'
	// 'MARKDOWN'
	// 常量不在本模块定义
	BodyType string `bson:"body_type"`
	// sender is user的uid ! 不是 contactinfo的uid
	Sender string `bson:"sender"`
	// 具体这个链接发布到了什么平台, youtube 还是 bilibili 直接通过url来识别, 不存储成 []{k:string, v:string} 的结构
	Link []string `bson:"link"`
	Tag  []string `bson:"tag"`
}

type ContentCommentModel struct {
	ID         string     `json:"_id" bson:"_id,omitempty"` // omitempty，务必
	CreatedAt  time.Time  `bson:"created_at"`               // service generate
	UpdatedAt  time.Time  `bson:"updated_at"`               // service generate
	DeletedAt  *time.Time `bson:"deleted_at"`               // service generate
	UID        string     `bson:"uid"`                      // service generate
	ContentUID string     `bson:"content_uid"`
	Title      string     `bson:"title"`
	Content    string     `bson:"content"`
	// 24小时之内未审核, 以及24小时以上已审核的内容会被显示
	Verified bool `bson:"verified"`
	// created_time将会跟article的updatetime相关
	// 如果article update, 并且
	Contact   UserContactModel                   `bson:"contact"`
	Reference ContentCommentModelExtra_Reference `bson:"reference"`
}

type ContentCommentModelExtra_Reference struct {
	BodyIndex    string `bson:"body_index"`
	BodyPartFrom int    `bson:"body_part_from"`
	BodyPart     string `bson:"body_part"`
}
