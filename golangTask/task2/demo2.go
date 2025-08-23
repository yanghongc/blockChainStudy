package main

import "fmt"

//题目 ：定义一个 Shape 接口，包含 Area() 和 Perimeter() 两个方法。
// 然后创建 Rectangle 和 Circle 结构体，实现 Shape 接口。在主函数中，
// 创建这两个结构体的实例，并调用它们的 Area() 和 Perimeter() 方法。
type Shape interface {
	Area()

	Perimeter()
}

type Rectangle struct {
	Name string
}

type Circle struct {
	Name string
}

func (r Rectangle) Area() {

	fmt.Printf("%v调用了Area方法", r.Name)

}

func (r Rectangle) Perimeter() {
	fmt.Printf("%v调用了Perimeter方法", r.Name)
}

func (c Circle) Area() {
	fmt.Printf("%v调用了Area方法", c.Name)
}

func (c Circle) Perimeter() {
	fmt.Printf("%v调用了Perimeter方法", c.Name)
}

func test6() {
	rec := Rectangle{
		Name: "Rectangle",
	}

	var sha1 Shape = rec //实现Shape接口
	sha1.Area()
	sha1.Perimeter()

	cir := Circle{
		Name: "Circle",
	}

	var sha2 Shape = cir
	sha2.Area()
	sha2.Perimeter()
}

//题目 ：使用组合的方式创建一个 Person 结构体，包含 Name 和 Age 字段，再创建一个 Employee 结构体，
// 组合 Person 结构体并添加 EmployeeID 字段。为 Employee 结构体实现一个 PrintInfo() 方法，输出员工的信息。
type Employee struct {
	per        Person
	EmployeeID string
}

type Person struct {
	Name string
	Age  int
}

func (e Employee) PrintInfo() {
	fmt.Printf("员工信息:\n姓名: %s\n年龄: %d\n工号: %s\n",
		e.per.Name, e.per.Age, e.EmployeeID)

}

func test7() {
	emp := Employee{
		per: Person{
			Name: "李四",
			Age:  28,
		},
		EmployeeID: "101",
	}
	emp.PrintInfo()
}
