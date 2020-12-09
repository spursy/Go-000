package controller

import (
	"github.com/gin-gonic/gin"
)

func PingFunc(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}