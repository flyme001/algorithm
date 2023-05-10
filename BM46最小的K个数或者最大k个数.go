描述
给定一个长度为 n 的可能有重复值的数组，找出其中不去重的最小的 k 个数。例如数组元素是4,5,1,6,2,7,3,8这8个数字，则最小的4个数字是1,2,3,4(任意顺序皆可)。
输入：
[4,5,1,6,2,7,3,8],4 
复制
返回值：
[1,2,3,4]
复制
说明：
返回最小的4个数即可，返回[1,3,2,4]也可以    

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
