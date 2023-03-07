package main

import "sort"

func minMoney(arr []int, money int) int {
	var m int

	index := make(map[int]int, 0)
	for k, v := range arr {
		index[v] = k
	}

	sort.Ints(arr)

	if money/arr[0] > 0 {
		if money%arr[0] == 0 {
			m = money / arr[0]
		}
	}

	temp := money - arr[0]
	for i := 1; i < len(arr); i++ {
		if temp == arr[i] {
			m = m + 1
		} else {
			if _, ok := index[temp-arr[i]]; ok {
				m = m + 2
			}
		}
	}

	return m
}

func main() {

}
