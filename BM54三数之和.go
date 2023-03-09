// 给出一个有n个元素的数组S，S中是否有元素a,b,c满足a+b+c=0？找出数组S中所有满足条件的三元组。
package main

import "fmt"

func threeSum(nums []int, target int) [][]int {
	res := [][]int{}
	l := len(nums) - 2

	for j := 0; j < l; j++ {
		arr := []int{}
		for i := j + 1; i < len(nums)-1; i++ {
			arr := append(arr, nums[j], nums[i], nums[i+1])
			sum := isZero(arr)
			if sum == 0 {
				res = append(res, arr)
			}
		}
	}

	return res
}

func isZero(num []int) int {
	sum := 0
	for i := 0; i < len(num); i++ {
		sum = sum + num[i]
	}
	return sum
}

func main() {
	nums := []int{3, 5, 6, -3, 0}
	ar := threeSum(nums, 0)
	fmt.Println(ar)
}
