package net

import "github.com/gin-gonic/gin"

var _netLint *gin.Engine

func Init_Login(netLink *gin.Engine) {
	_netLint = netLink

	_netLint.GET("/", func(c *gin.Context) {
		c.String(200, "Hello, Geektutu")
	})
}
