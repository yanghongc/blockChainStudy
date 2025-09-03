package modules

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"unique;not null" form:"username"`
	Password string `gorm:"not null" form:"password"`
	Email    string `gorm:"unique;not null" form:"email"`
}

type Post struct {
	gorm.Model
	Title   string `gorm:"not null" form:"title"`
	Content string `gorm:"not null" form:"content"`
	UserID  uint
	User    User
}

type Comment struct {
	gorm.Model
	Content string `gorm:"not null"`
	UserID  uint
	User    User
	PostID  uint
	Post    Post
}

var DB *gorm.DB
var err error

func init() {
	dsn := "root:Passwd@123@tcp(117.72.97.42:3306)/test?charset=utf8mb4&parseTime=True"

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(fmt.Errorf("failed to connect database: %w", err))
	}

	DB.AutoMigrate(&User{}, &Post{}, &Comment{})
}
