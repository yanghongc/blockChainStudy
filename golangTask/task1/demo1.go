package main

import "fmt"

/**
* 136. 只出现一次的数字 给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现两次。
* 找出那个只出现了一次的元素。可以使用 for 循环遍历数组，结合 if 条件判断和 map 数据结构来解决，
* 例如通过 map 记录每个元素出现的次数，然后再遍历 map 找到出现次数为1的元素。
 */
func test1() {
	//定义字符串
	var a = [9]int{2, 2, 3, 4, 5, 6, 4, 6, 5}
	//声明map
	var m = make(map[int]int)
	//for循环获取每个数字出现的次数
	for _, v := range a {
		m[v]++
	}
	fmt.Printf("每个值出现的次数：%v\n", m)
	//找值为1对应的k
	for k, v := range m {
		if v == 1 {
			fmt.Printf("只出现一次的数字：%v\n", k)
			break
		}
	}

}

/**
* 异或运算，相同为0，不同1
 */
func test2() {
	//定义字符串
	var a = [9]int{2, 2, 3, 4, 5, 6, 4, 6, 5}
	var result = 0
	for _, v := range a {
		result ^= v
	}
	fmt.Printf("只出现一次的数字：%v\n", result)
}

//回文数
func test3() {
	var a = 1234321
	var tmp = a
	var recver = 0
	for {
		recver = recver*10 + a%10
		fmt.Printf("倒序的数：%v\n", recver)
		a = a / 10
		if a == 0 {
			break
		}
	}

	if tmp == recver {
		fmt.Printf("%v是一个回文数", tmp)
	}

}

func main() {
	//test1() //只出现一次的数字
	//test2() //只出现一次的数字
	//test3() //回文数
	//test4() //有效的括号
	//test5() //公共前缀
	//test6()
	//test7()
	//test8()
	test9()
}
