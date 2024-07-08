package main

import (
	"NResoureceLibraryServer/db"
	"NResoureceLibraryServer/db/gorm"
	_ "NResoureceLibraryServer/net"

	"github.com/gin-gonic/gin"
)

func main() {

	LinkDB()

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(200, "Hello, Geektutu")
	})

	r.Run(":3100") // listen and serve on 0.0.0.0:8080
}

func LinkDB() {
	var dbConfig *gorm.DBConfig = new(gorm.DBConfig)
	dbConfig.DBName = "NResoureceLibraryServerDB"
	dbConfig.Host = "guiyewyc.top"
	dbConfig.Password = "159357"
	dbConfig.Port = "3306"
	dbConfig.UserName = "root"
	dbConfig.Type = gorm.DB_Mysql
	//链接数据库
	db.Init_DB(dbConfig)

}
