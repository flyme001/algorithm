package main

import "fmt"

// 冒泡算法
func bubbleSort(arr []int) {
	l := len(arr)

	for i := 0; i < l; i++ {
		for j := 0; j < l-1-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}

	fmt.Println(arr)
}

func maxValue(arr []int) int {
	max := 0
	if len(arr) == 0 {
		return 0
	}
	if len(arr) == 1 {
		return arr[0]
	}

	for i := 0; i < len(arr); i++ {
		if arr[i] > max {
			max = arr[i]
		}
	}
	fmt.Println(max)
	return max
}

func main() {
	arr := []int{1, 3, 5, 3, 4}
	bubbleSort(arr)
	maxValue(arr)
}
