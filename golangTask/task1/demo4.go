package main

import (
	"fmt"
	"sort"
)

/*
删除有序数组中的重复项: 给你一个有序数组 nums ，请你原地删除重复出现的元素，使每个元素只出现一次，
返回删除后数组的新长度。不要使用额外的数组空间，你必须在原地修改输入数组并在使用 O(1) 额外空间的条件下完成。
可以使用双指针法，一个慢指针 i 用于记录不重复元素的位置，一个快指针 j 用于遍历数组，当 nums[i] 与 nums[j] 不相等时，
将 nums[j] 赋值给 nums[i + 1]，并将 i 后移一位。

*/

func test7() {
	var nums = [6]int{1, 7, 6, 6, 6, 7}
	var i = 0
	for j := 1; j < len(nums); j++ {

		if nums[i] != nums[j] {
			nums[i+1] = nums[j]
			i++
		}
	}

	fmt.Printf("删除有序数组中的重复项后得到的数组：%v", nums[:i+1])
}

/*
合并区间：以数组 intervals 表示若干个区间的集合，其中单个区间为 intervals[i] = [starti, endi] 。
请你合并所有重叠的区间，并返回一个不重叠的区间数组，该数组需恰好覆盖输入中的所有区间。
可以先对区间数组按照区间的起始位置进行排序，然后使用一个切片来存储合并后的区间，遍历排序后的区间数组，
将当前区间与切片中最后一个区间进行比较，如果有重叠，则合并区间；如果没有重叠，则将当前区间添加到切片中。
*/
func test8() {
	var intervals = [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}
	// 按起始位置升序排序
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] == intervals[j][0] {
			return intervals[i][1] < intervals[j][1]
		}
		return intervals[i][0] < intervals[j][0]
	})

	res := make([][]int, 0, len(intervals))
	for _, iv := range intervals {
		start, end := iv[0], iv[1]
		// 如果结果为空或当前区间与最后一个区间不重叠，直接追加
		if len(res) == 0 || start > res[len(res)-1][1] {
			res = append(res, []int{start, end})
		} else {
			// 有重叠，合并：
			if end > res[len(res)-1][1] {
				res[len(res)-1][1] = end
			}
		}
	}
	fmt.Printf("合并后的区间：%v", res)

}

//给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那两个整数

func test9() {
	//整数数组
	var nums = [...]int{1, 3, 5, 7, 9}
	//目标值
	var target = 8
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				fmt.Printf("该数组中和为%v的那两个整数：{%v,%v}", target, nums[i], nums[j])
			}
		}
	}
}
