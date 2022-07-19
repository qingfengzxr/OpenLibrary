package model

import (
	"time"
)

// Book [...]
type Book struct {
	Cid      string    `gorm:"primaryKey;column:cid;type:varchar(100);not null" json:"-"` // cid
	Creator  string    `gorm:"column:creator;type:varchar(100);not null" json:"creator"`  // 创建者
	Chapter  int       `gorm:"column:chapter;type:int;not null;default:1" json:"chapter"` // 卷
	Section  int       `gorm:"column:section;type:int;not null" json:"section"`           // 章节
	Page     int       `gorm:"column:page;type:int;not null" json:"page"`                 // 页数
	Content  []byte    `gorm:"column:content;type:longblob;not null" json:"content"`      // 内容
	CreateAt time.Time `gorm:"column:createAt;type:datetime;not null" json:"create_at"`   // 创建时间
	UpdateAt time.Time `gorm:"column:updateAt;type:datetime;not null" json:"update_at"`   // 更新时间
	DeleteAt time.Time `gorm:"column:deleteAt;type:datetime;not null" json:"delete_at"`   // 删除时间
}

// TableName get sql table name.获取数据库表名
func (m *Book) TableName() string {
	return "book"
}

// BookColumns get sql column name.获取数据库列名
var BookColumns = struct {
	Cid      string
	Creator  string
	Chapter  string
	Section  string
	Page     string
	Content  string
	CreateAt string
	UpdateAt string
	DeleteAt string
}{
	Cid:      "cid",
	Creator:  "creator",
	Chapter:  "chapter",
	Section:  "section",
	Page:     "page",
	Content:  "content",
	CreateAt: "createAt",
	UpdateAt: "updateAt",
	DeleteAt: "deleteAt",
}
