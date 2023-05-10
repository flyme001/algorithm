给定一个无重复元素的整数数组nums，请你找出其中没有出现的最小的正整数
输入：
[1,0,2]
复制
返回值：
3
package main

import "sort"

func minNumberDisappeared(nums []int) int {
	sort.Ints(nums)

	index := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			index = index + 1
		}
		
		if nums[i] != index {
			return nums[i]
		}
	}

	return -1
}
