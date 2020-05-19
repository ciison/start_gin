package middlewares

import (
	"net/http"
	
	"github.com/gin-gonic/gin"
)

func MiddlewareCros(ctx *gin.Context) {
	
	ctx.Header("Access-Control-Allow-Origin", "*")
	ctx.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
	
	if http.MethodOptions == ctx.Request.Method {
		ctx.AbortWithStatusJSON(http.StatusOK, "Options Request!")
		return
	}
	
	ctx.Next()
	
}
