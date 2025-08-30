package main

import (
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql" // 不要忘了导入数据库驱动
	"github.com/jmoiron/sqlx"
)

/*
	题目1：使用SQL扩展库进行查询
	假设你已经使用Sqlx连接到一个数据库，并且有一个 employees 表，包含字段 id 、 name 、 department 、 salary 。
	要求 ：
	编写Go代码，使用Sqlx查询 employees 表中所有部门为 "技术部" 的员工信息，并将结果映射到一个自定义的 Employee 结构体切片中。
	编写Go代码，使用Sqlx查询 employees 表中工资最高的员工信息，并将结果映射到一个 Employee 结构体中。
*/

var db *sqlx.DB

func initDB() (err error) {
	dsn := "root:Passwd@123@tcp(117.72.97.42:3306)/test?charset=utf8mb4&parseTime=True"
	// 也可以使用MustConnect连接不成功就panic
	db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		fmt.Printf("connect DB failed, err:%v\n", err)
		return
	}
	db.SetMaxOpenConns(20)
	db.SetMaxIdleConns(10)
	return
}

type Employee struct {
	Id         int       `db:"id"`
	Name       string    `db:"name"`
	Department string    `db:"department"`
	Salary     float64   `db:"salary"`
	CreatedAt  time.Time `db:"created_at"`
	UpdatedAt  time.Time `db:"updated_at"`
}

// 编写Go代码，使用Sqlx查询 employees 表中所有部门为 "技术部" 的员工信息，并将结果映射到一个自定义的 Employee 结构体切片中。
var employees []Employee

func QueryInfo() {
	var sql = "select * from employees where department = ?"
	err := db.Select(&employees, sql, "技术部")
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
		return
	}
	fmt.Printf("employees:%#v\n", employees)
}

// 编写Go代码，使用Sqlx查询 employees 表中工资最高的员工信息，并将结果映射到一个 Employee 结构体中。
func QueryInfoBySal() {
	var sql = "select * from employees where salary = (select max(salary) from employees)"
	err := db.Select(&employees, sql)
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
		return
	}
	fmt.Printf("employees:%#v\n", employees)
}

/*
	题目2：实现类型安全映射
	假设有一个 books 表，包含字段 id 、 title 、 author 、 price 。
	要求 ：
	定义一个 Book 结构体，包含与 books 表对应的字段。
	编写Go代码，使用Sqlx执行一个复杂的查询，例如查询价格大于 50 元的书籍，并将结果映射到 Book 结构体切片中，确保类型安全。
*/

type Book struct {
	Id        int       `db:"id"`
	Title     string    `db:"title"`
	Author    string    `db:"author"`
	Price     float64   `db:"price"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

var books []Book

// 编写Go代码，使用Sqlx执行一个复杂的查询，例如查询价格大于 50 元的书籍，并将结果映射到 Book 结构体切片中，确保类型安全。
func QueryInfoByPrice() {
	var sql = "select * from books where price > ?"
	err := db.Select(&books, sql, 50)
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
		return
	}
	fmt.Printf("books:%#v\n", books)
}

func main() {
	//initDB()
	//QueryInfo()
	//QueryInfoBySal()

	//initDbGrom()

	//QueryInfoById(1)

	QueryInfoByCont()
}
