package mysqlDB

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql" //导入包但不使用，init()
)

var _dbName string
var _host string
var _userName string
var _password string
var db *sql.DB

// 初始化数据库
func InitMySqlDB(userName string, password string, host string, dbName string) {
	_dbName = dbName
	_password = password
	_host = host
	_userName = userName

	LinkedDatabase()
}

func LinkedDatabase() {

	//用户名:密码啊@tcp(ip:端口)/数据库的名字
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s", _userName, _password, _host, _dbName)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		fmt.Printf("dsn:%s invalid,err:%v\n", dsn, err)
		return
	}
	err = db.Ping() //尝试连接数据库
	if err != nil {
		fmt.Printf("open %s faild,err:%v\n", dsn, err)
		return
	}
	fmt.Println("连接数据库成功~")
}
