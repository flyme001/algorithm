package main

import (
	"sort"
)

package main

func maxProfit4(num []int) int {
	max := 0
	slow := 0
	arr := []int{}

	for i := 1; i < len(num); i++ {
		if num[i] > num[slow] {
			if num[i]-num[slow] > max {
				arr = append(arr, max + num[i] - num[slow])
			}
		} else {
			slow = i
		}
	}

	sort.Ints(arr)
	return max
}

func main() {
	maxProfit4([]int{8, 9, 2, 5, 4, 7, 1})
}
