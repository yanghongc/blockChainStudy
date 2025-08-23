package main

import "fmt"

//给定一个由整数组成的非空数组所表示的非负整数，在该数的基础上加一
func test6() {
	var numInt = []int{9, 9, 9, 9, 9}

	for i := len(numInt) - 1; i >= 0; i-- {
		if numInt[i] == 9 {
			numInt[i] = 0

		} else {
			numInt[i] += 1
			fmt.Printf("大整数的整数数组 :%v", numInt)
			break
		}
	}
	//当所有位数都是9的时候，前面插1
	if numInt[0] == 0 {
		fmt.Printf("大整数的整数数组 :%v", append([]int{1}, numInt...))
	}
}
