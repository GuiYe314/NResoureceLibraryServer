package main

import (
	"gee"
	"mysqlDB"
)

func main() {
	r := gee.New()
	//初始化数据库
	mysqlDB.InitMySqlDB("root", "159357", "guiyewyc.top:3306", "NResoureceLibraryServerDB")

	//开启监听
	gee.InitHttp(r)

	r.Run(":2080")
}
