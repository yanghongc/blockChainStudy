package main

import "fmt"

//给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效
func test4() {
	var str = "({[([(){}])]}{})"
	// 映射右括号 -> 左括号
	pairs := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}

	// 用 rune slice 作为栈
	var stack []rune

	for _, ch := range str {
		switch ch {
		// 左括号入栈
		case '(', '[', '{':
			stack = append(stack, ch)
		// 处理右括号
		case ')', ']', '}':
			// 如果栈为空或栈顶不匹配 -> 无效
			if len(stack) == 0 || stack[len(stack)-1] != pairs[ch] {
				break
			}
			// 弹出栈顶
			stack = stack[:len(stack)-1]
		}
	}
	// 栈为空则全部匹配
	if len(stack) == 0 {
		fmt.Println("该字符串为有效字符串")
	} else {
		fmt.Println("该字符串为无效字符串")
	}
}

//查找字符串数组中的最长公共前缀
func test5() {
	var str = [5]string{"abcoqwer", "abco", "abcyhiirgh", "abcvgyut", "abcohtgfdd"}
	//将第一个字符串作为公共前缀
	var comm = str[0]

	for i := 1; i < len(str); i++ {

		for j := 0; j < len(comm); j++ {
			//校验str[i]字符串是否有第j个字符
			if j >= len(str[i]) {

				comm = comm[:j]
				break
			}

			if comm[j] != str[i][j] {
				comm = comm[:j]
				fmt.Println("公共前缀是：" + comm)
				break
			}
		}
	}

}
