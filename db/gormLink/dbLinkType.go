package gormLink

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// mysql 结构体
var mySqlLinkTyep *MySqlLinkTyep

type MySqlLinkTyep struct {
	db *gorm.DB
}

func LinkDBType(dbCinfg *DBConfig) (*gorm.DB, error) {

	switch dbCinfg.Type {
	case DB_Mysql:
		mySqlLinkTyep = new(MySqlLinkTyep)
		return mySqlLinkTyep.LinkDB(dbCinfg)
	//case "postgres":
	default:
	}

	return nil, nil
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
	return db, nil
}
