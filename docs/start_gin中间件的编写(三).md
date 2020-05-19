# start_gin 中间件的编写

**gin的业务处理都是基于中间的的方式** `type HandlerFunc func(*Context)` 签名

所以, 只要参数是 `*gin.Context` 的函数都可以作为 中间件使用

以 跨域请求为例写一个中间件:

```go

func MiddlewareCros(ctx *gin.Context) {
	
	ctx.Header("Access-Control-Allow-Origin", "*")
	ctx.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
	
	if http.MethodOptions == ctx.Request.Method {
		ctx.AbortWithStatusJSON(http.StatusOK, "Options Request!")
		return
	}
	
	ctx.Next()
	
}
```

[完整代码](https://github.com/ciison/start_gin/tree/%E4%B8%AD%E9%97%B4%E4%BB%B6fixed)