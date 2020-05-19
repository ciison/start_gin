# start_gin 编写要给简单的业务


[参考文档](https://gin-gonic.com/zh-cn/docs/examples/http-method/)
```go

// 监听 /index 路由的 GET 请求
engine.GET("/index", func(ctx *gin.Context) {
    // 这里是将整个调用链设置到终止的那一层, 返回客户端 JSON 格式的数据
    ctx.AbortWithStatusJSON(200, gin.H{"status": "ok"})
})


```


同理 POST 请求

```go
engine.POST("/index", func(ctx *gin.Context) {
    ctx.AbortWithStatusJSON(200, gin.H{"METHOD": "post"})
})

```

还可以将参数绑定到 指定的struct

[我文档地址](https://gin-gonic.com/zh-cn/docs/examples/bind-body-into-dirrerent-structs/)