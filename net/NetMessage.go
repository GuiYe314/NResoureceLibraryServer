package net

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SendMassage(c *gin.Context, code int16, msg string) {
	c.JSON(http.StatusUnprocessableEntity, gin.H{
		"code":    code,
		"message": msg,
	})
}
