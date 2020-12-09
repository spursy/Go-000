package main

import (
	"github.com/Go-000/Week02/001homework/controller"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", controller.PingFunc)

	_ = r.Run(":8090")
}