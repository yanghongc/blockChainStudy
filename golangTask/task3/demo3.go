package main

import (
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// User 表示用户
type User struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time

	// 用户字段
	Name     string `gorm:"type:varchar(100);not null"`
	Email    string `gorm:"type:varchar(100);uniqueIndex;not null"`
	Password string `gorm:"type:varchar(255);not null"`

	PostCount int64 `gorm:"column:post_count;default:0"` // 新增字段：用户的文章数量统计
	// 关系：一个用户有多篇文章
	Posts []Post `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

// Post 表示文章
type Post struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time

	Title   string `gorm:"type:varchar(200);not null"`
	Content string `gorm:"type:text;not null"`

	// 外键：作者
	UserID uint
	User   User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`

	CommentStatus string `gorm:"column:comment_status;type:varchar(32);default:'无评论'"` // "有评论"/"无评论"
	// 关系：一篇文章有多条评论
	Comments []Comment `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

// Comment 表示评论
type Comment struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time

	Content string `gorm:"type:text;not null"`

	// 外键：所属文章
	PostID uint
	Post   Post `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

var gromDb *gorm.DB

var err error

func init() {
	dsn := "root:Passwd@123@tcp(117.72.97.42:3306)/test?charset=utf8mb4&parseTime=True"

	gromDb, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(fmt.Errorf("failed to connect database: %w", err))
	}

}

func initDbGrom() {

	crerr := gromDb.AutoMigrate(&User{}, &Post{}, &Comment{})

	if crerr != nil {
		panic(fmt.Errorf("auto migrate failed: %w", err))
	} else {
		fmt.Println("创建成功")
	}

}

/*
	题目2：关联查询
	基于上述博客系统的模型定义。
	要求 ：
	编写Go代码，使用Gorm查询某个用户发布的所有文章及其对应的评论信息。
	编写Go代码，使用Gorm查询评论数量最多的文章信息。
*/

// 编写Go代码，使用Gorm查询某个用户发布的所有文章及其对应的评论信息。
func QueryInfoById(userId int) {
	var userWithPosts User
	err := gromDb.Preload("Posts.Comments").First(&userWithPosts, userId).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			fmt.Printf("用户 %d 未找到\n", userId)
		} else {
			fmt.Printf("查询出错: %v\n", err)
		}
		return
	}

	fmt.Printf("用户: %s (ID=%d)\n", userWithPosts.Name, userWithPosts.ID)
	for _, p := range userWithPosts.Posts {
		fmt.Printf("  文章: %s (ID=%d)\n", p.Title, p.ID)
		for _, c := range p.Comments {
			fmt.Printf("    评论: %s (ID=%d, 作者=%s)\n", c.Content)
		}
	}
}

// 编写Go代码，使用Gorm查询评论数量最多的文章信息。
type PostWithCount struct {
	ID           uint
	Title        string
	Content      string
	UserID       uint
	CommentCount int64 `gorm:"column:comment_count"`
}

func QueryInfoByCont() {
	// 查询前 N 篇按评论数降序的文章（例如前 5）
	limit := 5
	var results []PostWithCount

	// 方法：从 posts 表出发，LEFT JOIN comments 统计数量，并 JOIN users 获取作者名
	err = gromDb.
		Table("posts").
		Select("posts.id, posts.title, posts.content, posts.user_id, posts.created_at, users.name AS author_name, COUNT(comments.id) AS comment_count").
		Joins("LEFT JOIN comments ON comments.post_id = posts.id").
		Group("posts.id").
		Order("comment_count DESC").
		Limit(limit).
		Scan(&results).Error
	if err != nil {
		fmt.Printf("未找到")
	}

	fmt.Printf("查询出的返回结果：%v\n", results)
}

/*
	题目3：钩子函数
	继续使用博客系统的模型。
	要求 ：
	为 Post 模型添加一个钩子函数，在文章创建时自动更新用户的文章数量统计字段。
	为 Comment 模型添加一个钩子函数，在评论删除时检查文章的评论数量，如果评论数量为 0，则更新文章的评论状态为 "无评论"。
*/
// Post 模型方法：创建后增加用户文章计数
func (p *Post) AfterCreate() {
	// 在数据库层做自增，避免读取-修改-写入的竞态
	err := gromDb.Model(&User{}).
		Where("id = ?", p.UserID).
		UpdateColumn("post_count", gorm.Expr("post_count + ?", 1)).Error

	if err != nil {
		fmt.Printf("创建后增加用户文章计数出错: %v\n", err)
	}
}

// 删除Comment后回调，检查评论状态
func (c *Comment) AfterDelete() {

	// 获取该Post的当前评论数量
	var count int64
	gromDb.Model(&Post{}).Where("id = ?", c.PostID).Count(&count)

	if count == 0 {
		gromDb.Model(&Post{}).Where("id = ?", c.PostID).Update("comment_status", "无评论")
	}

}
