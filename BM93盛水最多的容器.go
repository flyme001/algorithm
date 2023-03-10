package main

import "sort"

func moreWater(num []int) int {
	sort.Ints(num)
	for i := 0; i < len(num)-1; i++ {
		if num[len(num)-1-i] == num[len(num)-2-i] {
			return num[len(num)-1-i] * num[len(num)-1-i]
		}
	}
	return 0
}
