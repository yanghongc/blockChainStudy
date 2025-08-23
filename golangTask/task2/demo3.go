package main

import (
	"fmt"
)

// 题目 ：编写一个程序，使用通道实现两个协程之间的通信。
// 一个协程生成从1到10的整数，并将这些整数发送到通道中，
// 另一个协程从通道中接收这些整数并打印出来。

func writeData(inChain chan int) {
	for i := 0; i < 10; i++ {
		inChain <- i
		fmt.Println("writeData 写入数据-", i+1)
	}
	close(inChain)
	wg.Done()

}

func readData(inChain chan int) {
	for v := range inChain {
		fmt.Println("readData 读取数据-", v)
	}

	wg.Done()
}

func test8() {
	allChain := make(chan int, 10)
	wg.Add(1)
	go writeData(allChain)

	wg.Add(1)
	go readData(allChain)

	wg.Wait()

	fmt.Println("读取完毕。。")

}
