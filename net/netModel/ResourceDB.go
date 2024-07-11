package netModel

import (
	"database/sql"
	"time"
)

// 目录数据表
type Contents struct {
	ID            uint   `gorm:"primarykey"` //ID
	Parent_id     uint   //父ID
	Contents_Name string `gorm:"size:255"` //目录名
	Grade         uint   //所在层级
	Created_At    time.Time
	Updated_At    time.Time
	Deleted_At    sql.NullTime `gorm:"index"`
}

type Tags struct {
	ID         uint   `gorm:"primarykey"`
	Tags_Name  string `gorm:"size:255"` //标签名
	Type       string `gorm:"size:255"` //标签类型
	Created_At time.Time
	Updated_At time.Time
	Deleted_At sql.NullTime `gorm:"index"`
}

type TagsType struct {
	ID            uint   `gorm:"primarykey"`
	TagsType_Name string `gorm:"size:255"` //标签名
	Created_At    time.Time
	Updated_At    time.Time
	Deleted_At    sql.NullTime `gorm:"index"`
}
