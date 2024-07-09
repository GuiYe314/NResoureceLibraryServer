package netModel

import (
	"database/sql"
	"time"
)

// Users 结构体对应users表中的一行数据
type Users struct {
	ID         uint   `gorm:"primarykey"`
	User_Name  string `gorm:"size:255;unique"` // 用户名，唯一
	Password   string `gorm:"size:255"`        // 密码
	Email      string `gorm:"size:255;unique"` // 电子邮件，唯一
	Gender     string `gorm:"size:10"`         // 性别
	Created_At time.Time
	Updated_At time.Time
	Deleted_At sql.NullTime `gorm:"index"`
}
