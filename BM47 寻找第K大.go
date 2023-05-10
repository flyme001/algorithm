有一个整数数组，请你根据快速排序的思路，找出数组中第 k 大的数。

给定一个整数数组 a ,同时给定它的大小n和要找的 k ，请返回第 k 大的数(包括重复的元素，不用去重)，保证答案存在。
输入：
[1,3,5,2,2],5,3
复制
返回值：
2

package main

import "sort"

// 数组进行排序 然后取k下标值
func ksort(arr []int, k int) int {
	sort.Ints(arr)

	return arr[k-1]
}

func main() {
}
