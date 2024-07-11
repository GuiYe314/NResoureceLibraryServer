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

// 标签类型
type TagsType struct {
	ID            uint   `gorm:"primarykey"`
	Tags_ID       string `gorm:"size:255"` //标签ID
	Tagstype_Name string `gorm:"size:255"` //标签名
	Created_At    time.Time
	Updated_At    time.Time
	Deleted_At    sql.NullTime `gorm:"index"`
}

type Resource struct {
	ID            uint   `gorm:"primarykey"`
	Resource_Name string `gorm:"size:255"` //资源名
	Resource_Path string `gorm:"size:255"` //资源路径-相对路径
	Contents_ID   string `gorm:"size:255"` //目录ID
	Tags_ID       string `gorm:"size:255"` //标签ID以#号键隔开
	Created_At    time.Time
	Updated_At    time.Time
	Deleted_At    sql.NullTime `gorm:"index"`
}

type ResourceImage struct {
	ID          uint   `gorm:"primarykey"`
	Image_Name  string `gorm:"size:255"` //资源名
	Image_Path  string `gorm:"size:255"` //资源路径-相对路径
	Resource_ID string `gorm:"size:255"` //属于那个资源
	Created_At  time.Time
	Updated_At  time.Time
	Deleted_At  sql.NullTime `gorm:"index"`
}

type ResourceVideo struct {
	ID          uint   `gorm:"primarykey"`
	Video_Name  string `gorm:"size:255"` //资源名
	Video_Path  string `gorm:"size:255"` //资源路径-相对路径
	Resource_ID string `gorm:"size:255"` //属于那个资源
	Created_At  time.Time
	Updated_At  time.Time
	Deleted_At  sql.NullTime `gorm:"index"`
}

type ResourceModel struct {
	ID          uint   `gorm:"primarykey"`
	Model_Name  string `gorm:"size:255"` //资源名
	Model_Path  string `gorm:"size:255"` //资源路径-相对路径
	Resource_ID string `gorm:"size:255"` //属于那个资源
	Created_At  time.Time
	Updated_At  time.Time
	Deleted_At  sql.NullTime `gorm:"index"`
}
