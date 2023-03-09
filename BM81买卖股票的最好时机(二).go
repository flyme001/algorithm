package main

func maxProfit3(num []int) int {
	max := 0
	slow := 0

	for i := 1; i < len(num); i++ {
		if num[i] > num[slow] {
			if num[i]-num[slow] > max {
				max = max + num[i] - num[slow]
			}
		} else {
			slow = i
		}
	}

	return max
}

func main() {
	maxProfit3([]int{8, 9, 2, 5, 4, 7, 1})
}
