package main

import "sort"

// 数组进行排序 然后取k下标值
func ksort(arr []int, k int) int {
	sort.Ints(arr)

	return arr[k-1]
}

func main() {
}
