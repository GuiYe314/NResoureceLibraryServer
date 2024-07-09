package net

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Init(dbLink *gorm.DB, netLink *gin.Engine) {
	fmt.Printf("测试")
}
