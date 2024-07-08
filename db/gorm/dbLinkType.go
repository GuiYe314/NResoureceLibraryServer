package gorm

import (
	"database/sql"
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// mysql 结构体
var mySqlLinkTyep *MySqlLinkTyep

type MySqlLinkTyep struct {
	db *gorm.DB
}

func LinkDBType(dbCinfg *DBConfig) {

	switch dbCinfg.Type {
	case DB_Mysql:
		mySqlLinkTyep = new(MySqlLinkTyep)
		mySqlLinkTyep.LinkDB(dbCinfg)
	//case "postgres":
	default:
	}
}

// 连接类型
func (mySqlLinkTyep MySqlLinkTyep) LinkDB(dbConfig *DBConfig) (*gorm.DB, error) {

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%v)/%s?charset=utf8&parseTime=True&loc=Local", dbConfig.UserName, dbConfig.Password, dbConfig.Host, dbConfig.Port, dbConfig.DBName)

	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       dsn,   // DSN data source name
		DefaultStringSize:         256,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据当前 MySQL 版本自动配置
	}), &gorm.Config{})

	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	mySqlLinkTyep.db = db
	fmt.Println("数据库链接成功：" + db.Name())

	//测试创建表
	err = db.AutoMigrate(&User{})
	if err != nil {
		panic("无法创建users表")
	}

	return db, nil
}

// User 结构体对应users表中的一行数据
type User struct {
	ID        uint   `gorm:"primarykey"`
	UserID    string // 用户ID（通常与主键ID相同或使用UUID等唯一标识符）
	UserName  string `gorm:"size:255;unique"` // 用户名，唯一
	Password  string `gorm:"size:255"`        // 密码
	Email     string `gorm:"size:255;unique"` // 电子邮件，唯一
	Gender    string `gorm:"size:10"`         // 性别
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt sql.NullTime `gorm:"index"`
}
