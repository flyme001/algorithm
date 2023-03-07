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