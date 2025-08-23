package main

import (
	"fmt"
	"sync"
	"time"
)

// 题目 ：编写一个Go程序，定义一个函数，该函数接收一个整数指针作为参数，
// 在函数内部将该指针指向的值增加10，然后在主函数中调用该函数并输出修改后的值。
func test1(num *int) {
	*num += 10
}

// 题目 ：实现一个函数，接收一个整数切片的指针，将切片中的每个元素乘以2。
func test2(sliNum *[]int) {
	for i := 0; i < len(*sliNum); i++ {
		(*sliNum)[i] *= 2
	}
}

var wg sync.WaitGroup

// 题目 ：编写一个程序，使用 go 关键字启动两个协程，一个协程打印从1到10的奇数，另一个协程打印从2到10的偶数。
func test3() {
	var sliNum = make([]int, 0, 10)
	for i := 1; i <= 10; i++ {
		if i%2 != 0 {
			sliNum = append(sliNum, i)
		}
	}
	fmt.Printf("1-10中的奇数：%v", sliNum)
	wg.Done() // 4、goroutine 结束就登记-1
}

func test4() {
	var sliNum = make([]int, 0, 10)
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			sliNum = append(sliNum, i)
		}
	}
	fmt.Printf("1-10中的偶数：%v", sliNum)
	wg.Done() // 4、goroutine 结束就登记-1
}

// 题目 ：设计一个任务调度器，接收一组任务（可以用函数表示），并使用协程并发执行这些任务，同时统计每个任务的执行时间
func test5(n int) {
	start := time.Now().UnixMilli()
	for num := (n-1)*30000 + 1; num <= n*30000; num++ {
		for i := 2; i < num; i++ {
			if num%i == 0 {
				break
			}
		}
	}
	wg.Done()
	end := time.Now().UnixMilli()
	fmt.Println("普通的方法耗时=", end-start)
}

func main() {
	/*
		var num = 5
		test1(&num)
		fmt.Println("修改后的值:", num)
	*/

	/*
		var sliNum = []int{2, 3, 5, 7}
		test2(&sliNum)
		fmt.Println("修改后的值:", sliNum)
	*/

	/*
		wg.Add(1) //2、启动一个 goroutine 就登记+1
		go test3()
		wg.Add(1) //2、启动一个 goroutine 就登记+1
		go test4()
		wg.Wait() // 3、等待所有登记的 goroutine 都结束
	*/

	/*
		for i := 1; i <= 4; i++ {
			wg.Add(1)
			go test5(i)
		}
		wg.Wait()
	*/

	//test6()

	//test7()

	//test8()

	//test9()

	test10()

}
