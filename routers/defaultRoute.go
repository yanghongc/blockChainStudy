package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/yanghongc/blockChainStudy/controller/admin"
	"github.com/yanghongc/blockChainStudy/controller/defaultInit"
	"github.com/yanghongc/blockChainStudy/controller/tall"
)

// 注册和登录
func DefaultRoutersInit(r *gin.Engine) {
	defaultRouters := r.Group("/")
	{
		defaultRouters.GET("/", defaultInit.DefaultController{}.Index)
		defaultRouters.POST("/register", defaultInit.DefaultController{}.Register)
		defaultRouters.POST("/login", defaultInit.DefaultController{}.Login)

	}
}

/*
实现文章的创建功能，只有已认证的用户才能创建文章，创建文章时需要提供文章的标题和内容。
实现文章的读取功能，支持获取所有文章列表和单个文章的详细信息。
实现文章的更新功能，只有文章的作者才能更新自己的文章。
实现文章的删除功能，只有文章的作者才能删除自己的文章。
*/
func TitleConRoutersInit(r *gin.Engine) {
	//登录认证
	defaultRouters := r.Group("/admin", defaultInit.CheckLogin)
	{
		defaultRouters.POST("/", admin.PostController{}.Admin)
		//实现文章的创建功能，只有已认证的用户才能创建文章，创建文章时需要提供文章的标题和内容。
		defaultRouters.POST("/add", admin.PostController{}.AddCont)
		//实现文章的删除功能，只有文章的作者才能删除自己的文章。
		defaultRouters.POST("/del", admin.PostController{}.DelCont)
		//实现文章的更新功能，只有文章的作者才能更新自己的文章。
		defaultRouters.POST("/upd", admin.PostController{}.UpdCont)
		//实现文章的读取功能，支持获取所有文章列表和单个文章的详细信息。
		defaultRouters.GET("/sel", admin.PostController{}.SelCont)
	}
}

/*
实现评论的创建功能，已认证的用户可以对文章发表评论。
实现评论的读取功能，支持获取某篇文章的所有评论列表。
*/
func TallConRoutersInit(r *gin.Engine) {
	//登录认证
	defaultRouters := r.Group("/tall", defaultInit.CheckLogin)
	{
		//实现评论的创建功能，已认证的用户可以对文章发表评论。
		defaultRouters.POST("/add", tall.TallControllerInit{}.AddTall)
		//实现评论的读取功能，支持获取某篇文章的所有评论列表。
		defaultRouters.GET("/sel", tall.TallControllerInit{}.SelTall)
	}
}
