package main

import "sort"

func minknums(nums []int, k int) []int {
	sort.Ints(nums)

	res := []int{}
	for i := 0; i < k; i++ {
		res = append(res, nums[i])
	}
	return res
}

func maxknums(nums []int, k int) []int {
	sort.Ints(nums)

	res := []int{}
	for i := 0; i < k; i++ {
		res = append(res, nums[len(nums)-i])
	}

	return res
}

func main() {
	nums := []int{1, 4, 5, 6, 7, 4}
	minknums(nums, 3)
	maxknums(nums, 3)
}