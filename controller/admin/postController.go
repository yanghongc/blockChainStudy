package admin

import (
	"fmt"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/yanghongc/blockChainStudy/modules"
)

type PostController struct {
}

// 系统登录的首页
func (con PostController) Admin(c *gin.Context) {
	//设置sessions
	session := sessions.Default(c)
	//配置session的过期时间
	session.Options(sessions.Options{
		MaxAge: 3600 * 6, // 6hrs   MaxAge单位是秒
	})

	session.Save() //设置session的时候必须调用

	c.HTML(http.StatusOK, "admin/index.html", gin.H{})
}

func (pot PostController) AddCont(c *gin.Context) {
	var post modules.Post
	if err := c.ShouldBind(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	uid, exists := c.Get("user_id")
	if exists {
		u, _ := uid.(uint)
		post.UserID = u
	}

	result := modules.DB.Create(&post) // 通过数据的指针来创建
	if result.RowsAffected > 1 {
		fmt.Print(post)
	}
	c.String(http.StatusOK, "add 成功")
}

func (pot PostController) DelCont(c *gin.Context) {
	var post = modules.Post{
		Title: c.PostForm("title"),
	}

	result := modules.DB.Delete(&post) // 通过数据的指针来创建
	if result.RowsAffected > 1 {
		fmt.Print(post)
	}
	c.String(http.StatusOK, "del 成功")
}

func (pot PostController) SelCont(c *gin.Context) {
	post := []modules.Post{}
	modules.DB.Find(&post)
	c.JSON(http.StatusOK, gin.H{"success": true, "result": post})
}

func (pot PostController) UpdCont(c *gin.Context) {
	post := modules.Post{
		Title: c.PostForm("title"),
	}
	modules.DB.Find(&post)

	post.Content = c.PostForm("content")
	modules.DB.Save(post)

	c.String(http.StatusOK, "upd 成功")
}
