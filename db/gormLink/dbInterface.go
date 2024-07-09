package gormLink

import "gorm.io/gorm"

type DBConfig struct {
	IdentificationName string // IdentificationName 用于获取特定的数据库连接。
	DBName             string //数据库名字
	UserName           string //用户名
	Password           string //密码
	Host               string //地址
	Port               string //端口
	Type               string // Type of the database ("mysql", "postgres", "mssql", etc.).
	SSLMode            string
	TimeZone           string
	dialector          gorm.Dialector
}

const (
	DB_Mysql string = "MySql"
)

//接口
//连接数据库接口
type Link_DB interface {
	LinkDB(dbConfig *DBConfig) (*gorm.DB, error)
}
