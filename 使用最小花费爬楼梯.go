package main

import "math"

func minCostClimbingStairs(cost []int) int {
	if len(cost) == 0 {
		return -1
	}
	if len(cost) == 1 {
		return cost[0]
	}
	if len(cost) == 2 {
		return int(math.Min(float64(cost[0]), float64(cost[1])))
	}

	sum := 0
	for i := 0; i < len(cost); {
		if cost[i]+cost[i+1] > cost[i]+cost[i+2] {
			i = i + 2
		} else {
			i = i + 1
		}
		sum = sum + 1
	}
	return sum
}

func main() {
	cost := []int{1, 3, 4, 1, 5}
	minNumberDisappeared(cost)
}
