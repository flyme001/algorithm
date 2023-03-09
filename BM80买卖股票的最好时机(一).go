package main

import "fmt"

func maxProfit(num []int) int {
	max := make(map[string]int)

	for i := 0; i < len(max); i++ {
		for j := i + 1; j < len(max)-1; j++ {
			max["max"] = num[0]
			if num[j] > num[i] {
				if num[j]-num[i] > max["max"] {
					max["max"] = num[j] - num[i]
				}
			}
		}
	}
	return max["max"]
}

func maxProfit2(num []int) int {
	max := 0
	slow := 0
	for i := 1; i < len(num); i++ {
		if num[i] > num[slow] {
			if num[i]-num[slow] > max {
				max = num[i] - num[slow]
			}
		} else {
			slow = i
		}
		fmt.Println(i, "---", slow)

	}
	fmt.Println("--max--", max)
	return max
}

func main() {
	maxProfit2([]int{8, 9, 2, 5, 4, 7, 1})
}
