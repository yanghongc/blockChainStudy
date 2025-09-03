package defaultInit

import (
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/yanghongc/blockChainStudy/modules"
	"golang.org/x/crypto/bcrypt"
)

type DefaultController struct{}

// 系统登录的首页
func (con DefaultController) Index(c *gin.Context) {
	//设置sessions
	session := sessions.Default(c)
	//配置session的过期时间
	session.Options(sessions.Options{
		MaxAge: 3600 * 6, // 6hrs   MaxAge单位是秒
	})
	session.Set("username", "张三 111")
	session.Save() //设置session的时候必须调用

	c.HTML(http.StatusOK, "default/index.html", gin.H{
		"msg": "我是一个msg",
		"t":   1629788418,
	})
}

// 注册用户
func (con DefaultController) Register(c *gin.Context) {
	var user modules.User
	if err := c.ShouldBind(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// 加密密码
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}
	user.Password = string(hashedPassword)

	//存储注册的数据
	if err := modules.DB.Save(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error})
		return
	}
	c.HTML(http.StatusOK, "default/login.html", gin.H{
		"username": user.Username,
		"email":    user.Email,
	})
}

// 用户登录
func (con DefaultController) Login(c *gin.Context) {
	var user modules.User
	if err := c.ShouldBind(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var storedUser modules.User
	if err := modules.DB.Where("username = ?", user.Username).First(&storedUser).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}

	// 验证密码
	if err := bcrypt.CompareHashAndPassword([]byte(storedUser.Password), []byte(user.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}

	// 生成 JWT
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":       storedUser.ID,
		"username": storedUser.Username,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString([]byte("your_secret_key"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	//设置sessions
	session := sessions.Default(c)
	//配置session的过期时间
	session.Options(sessions.Options{
		MaxAge: 3600 * 6, // 6hrs   MaxAge单位是秒
	})

	// 仅保存非敏感字段
	session.Set("user_id", storedUser.ID)
	session.Set("username", storedUser.Username)
	session.Set("token", tokenString)

	session.Save() //设置session的时候必须调用

	c.HTML(http.StatusOK, "admin/index.html", gin.H{
		"username": storedUser.Username,
	})
}

// 校验登录
func CheckLogin(c *gin.Context) {

	session := sessions.Default(c)
	// 取 session 字段
	uid := session.Get("user_id")
	token := session.Get("token")
	username := session.Get("username")

	// 判断 session 是否为空或字段缺失
	if uid == nil || token == nil || username == nil {
		// 如果需要返回 JSON 错误
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":   "unauthorized",
			"message": "user not logged in or session expired",
		})
		// 阻止后续处理
		c.Abort()
		return
	}

	// 将用户信息放入上下文，方便后续 handler 使用
	c.Set("user_id", uid)
	c.Set("username", username)
	c.Set("token", token)

	// 放行请求继续处理
	c.Next()
}
