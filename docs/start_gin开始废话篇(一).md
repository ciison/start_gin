# start gin 开始废话篇, 项目搭建
> 参考 https://github.com/gin-gonic/gin 必看, 官方网站的文档有点跟不上 /(ㄒoㄒ)/~~

前提: go 版本要支持 go modules 
设置 go modules

```shell script
export GO111MODULE=on
export GOPROXY=https://goproxy.cn
```
初始化项目: 使用 go module 的方式
```shell script
go mod init start_gin 
```

编写 main.go 文件
```go
package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	gin.Default().Run()
}

```
然后在终端运行 
```shell script
go mod tidy # 获取依赖的包
```
运行 main.go 
```shell 
go run main.go 
```
在控制台会看到如下输出: 嗯, 这就是起来一个 http 服务了
```shell script
[GIN-debug] [WARNING] Creating an Engine instance with the Logger and Recovery middleware already attached.

[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:	export GIN_MODE=release
 - using code:	gin.SetMode(gin.ReleaseMode)

[GIN-debug] Environment variable PORT is undefined. Using port :8080 by default
[GIN-debug] Listening and serving HTTP on :8080
```

## gin.Default().Run() 是怎么跑起来的
* Default() 做了什么?
```go
// Default returns an Engine instance with the Logger and Recovery middleware already attached.
func Default() *Engine {
	debugPrintWARNINGDefault()
	engine := New()
	engine.Use(Logger(), Recovery())
	return engine
}

```

Default() 函数主要就是初始化 Engine 对象, 其中使用了连个中间件

1. Logger() gin 默认提供的 Log 中间件
2. Recovery()  panic 恢复中间件

* Run() 做了什么?
```go
// Run attaches the router to a http.Server and starts listening and serving HTTP requests.
// It is a shortcut for http.ListenAndServe(addr, router)
// Note: this method will block the calling goroutine indefinitely unless an error happens.
func (engine *Engine) Run(addr ...string) (err error) {
	defer func() { debugPrintError(err) }()

	address := resolveAddress(addr)
	debugPrint("Listening and serving HTTP on %s\n", address)
	err = http.ListenAndServe(address, engine)
	return
}
```

主要是 resolveAddress 和 http.ListenAndServe

resolveAddress 主要是解析监听的端口, 参数长度至多为一个, 如果没有传入, 默认监听 `:8080`

http.ListenAndServe 是标准库的提供的方法



## 如何定制 gin 的 engine 

```go
engine := gin.New() 
engine.Use( // 使用日志和 Recovery 的中间件
    gin.Logger(),
    gin.Recovery(),
)
```

还是使用其他的中间件, 比如要把日志写到 es 

```go
func logger2es() gin.HandlerFunc {
	// init es client 
	esClient:= fmt.Printf // 你们就当这个是 esClient 吧
	return func(ctx *gin.Context) {
		method := ctx.Request.Method
		remote := ctx.Request.RemoteAddr
		path := ctx.Request.RequestURI
		start := time.Now()
		// write log 
		ctx.Next()
		esClient("[info] method:%s, remote:%s, path:%s, lantency:%v\n", method, remote, path, time.Now().Sub(start))
	}
}
```

应用 写日志到 es 的中间件

```go
engine.Use(logger2es())
```

