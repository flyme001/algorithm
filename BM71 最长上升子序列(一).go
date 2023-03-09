package main

import "fmt"

func maxQ(arr []int) {
	res := []int{}

	for i := 0; i < len(arr)-1; i++ {
		j := i + 1
		fmt.Println(arr[i], "----", arr[j])

		if arr[j] < arr[i] {
			continue
		}

		if i == len(arr)-1-1 {
			if arr[len(arr)-1] > arr[len(arr)-2] {
				res = append(res, arr[len(arr)-1])
			}
		}

		res = append(res, arr[i])

	}
	fmt.Println(res)
}

func main() {
	maxQ([]int{6, 3, 1, 5, 2, 3, 7})
}
