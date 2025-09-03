<img width="869" height="276" alt="image" src="https://github.com/user-attachments/assets/cb813d89-f867-4d06-b4dd-b22b9f5f8d73" />基于 Go、Gin 和 GORM 开发的个人博客系统后端

功能特性 ：

1.用户注册与登录（JWT 认证）

2.文章的创建、读取、更新和删除（CRUD）

3.文章的评论功能

4.数据验证和错误处理

5.数据库关系和外键约束

6.请求日志记录

技术栈 编程语言: Go 1.23+

Web 框架: Gin

ORM: GORM

数据库: MySQL

认证: JWT (JSON Web Tokens)

密码加密: bcrypt

项目结构 :
   ── controllers/ # 控制器层 
   ── modules/ # 数据库连接，数据模型
   ── routes/ # 路由转发 
   ── security/ # 用户登录注册 
   ── template/ # 模版
   ── go.mod # Go 模块定义 
   ── go.sum # 依赖校验和 
   ── main.go # 应用入口

MySQL 8.0

Git

安装步骤

克隆项目 git clone -b blog https://github.com/yanghongc/blockChainStudy.git

数据库配置 


JWT 配置


服务器配置
SERVER_PORT=9000

用户注册 URL: POST /register
用户登录 URL: POST /login

文章端点：认证: 需要用户登录后才能操作

获取所有文章 URL: GET /admin/sel

删除文章 URL: POST /admin/del

创建文章 URL: POST /admin/add

修改文章 URL: POST /admin/upd


评论端点：认证: 需要用户登录后才能操作
获取文章评论 URL: GET /tall/sel/

创建评论 URL: POST /tall/add

