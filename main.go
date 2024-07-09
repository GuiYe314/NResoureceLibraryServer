package main

import (
	"NResoureceLibraryServer/db"
	"NResoureceLibraryServer/db/gormLink"
	"NResoureceLibraryServer/net"
	_ "NResoureceLibraryServer/net"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func main() {

	//连接数据库
	dbLink, err := LinkDB()
	if err != nil {
		return
	}

	//初始化接口
	r := gin.Default()
	net.Init(dbLink, r)

	r.Run(":3100") // listen and serve on 0.0.0.0:8080
}

func LinkDB() (*gorm.DB, error) {
	var dbConfig *gormLink.DBConfig = new(gormLink.DBConfig)
	dbConfig.DBName = "NResoureceLibraryServerDB"
	dbConfig.Host = "guiyewyc.top"
	dbConfig.Password = "159357"
	dbConfig.Port = "3306"
	dbConfig.UserName = "root"
	dbConfig.Type = gormLink.DB_Mysql
	//链接数据库
	return db.Init_DB(dbConfig)

}
