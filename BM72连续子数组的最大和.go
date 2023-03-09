package main

import "fmt"

func maxArr(arr []int) int {
	l := len(arr)
	max := 0

	for i := 0; i < l; i++ {
		for j := i + 1; j <= l; j++ {
			s := sum(arr[i:j])
			if s > max {
				max = s
			}
		}
	}
	fmt.Println("----max--", max)
	return max
}

func sum(num []int) int {
	fmt.Println("-----num--", num)
	s := 0
	for i := 0; i < len(num); i++ {
		s = s + num[i]
	}
	return s
}

func main() {
	maxArr([]int{1, -2, 3, 10, -4, 7, 2, -5})
}
