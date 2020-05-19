package main

import (
	"github.com/gin-gonic/gin"
	"start_gin/middlewares"
)

func main() {
	engine := gin.Default()
	
	engine.Use(middlewares.MiddlewareCros) // 注册使用中间件
	
	engine.GET("/index", func(ctx *gin.Context) {
		ctx.AbortWithStatusJSON(200, gin.H{"status": "ok"})
	})
	
	err := engine.Run()
	if err != nil {
		panic(err)
	}
}

