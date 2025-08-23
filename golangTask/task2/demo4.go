package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

//题目 ：编写一个程序，使用 sync.Mutex 来保护一个共享的计数器。
// 启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。

var mutex sync.Mutex

var count int64 = 0

func sum1() {

	defer wg.Done()

	for i := 0; i < 1000; i++ {
		mutex.Lock()
		count++
		mutex.Unlock()
	}
}

func test9() {
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go sum1()
	}
	wg.Wait()

	fmt.Printf("累计的数：%v", count)
}

//题目 ：使用原子操作（ sync/atomic 包）实现一个无锁的计数器。
// 启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。

func sum2() {

	defer wg.Done()

	for i := 0; i < 1000; i++ {

		atomic.AddInt64(&count, 1)

	}
}

func test10() {
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go sum2()
	}
	wg.Wait()

	fmt.Printf("累计的数：%v", count)
}
