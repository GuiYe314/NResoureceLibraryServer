package net

import (
	"NResoureceLibraryServer/net/netModel"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Init(dbLink *gorm.DB, netLink *gin.Engine) {

	//认证接口组
	auth := netLink.Group("/auth")

	auth.Use(jwtAuthMiddleware())
	{
		auth.GET("/FindModelAll", FindModelAll)
	}

	CreateDataTable(dbLink)
	Init_Login(dbLink, netLink)
}

// 自动创建数据库
func CreateDataTable(dbLink *gorm.DB) {
	dbLink.AutoMigrate(&netModel.Users{})
	dbLink.AutoMigrate(&netModel.Contents{})
}
