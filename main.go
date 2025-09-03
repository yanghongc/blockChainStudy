package main

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/yanghongc/blockChainStudy/routers"
)

func main() {
	//创建路由引擎
	r := gin.Default()

	//加载模版
	r.LoadHTMLGlob("templates/**/*")

	// 创建基于 cookie 的存储引擎，secret11111 参数是用于加密的密钥
	store := cookie.NewStore([]byte("secret111"))
	//配置session的中间件 store是前面创建的存储引擎，我们可以替换成其他存储引擎
	r.Use(sessions.Sessions("mysession", store))

	//配置路由
	routers.DefaultRoutersInit(r)

	routers.TitleConRoutersInit(r)

	//启动服务
	r.Run(":9000")
}
