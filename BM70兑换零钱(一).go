package main

import "sort"

func minMoney(arr []int, money int) int {
	var m int

	index := make(map[int]int, 0)
	for k, v := range arr {
		index[v] = k
	}

	sort.Ints(arr)

	if money/arr[0] > 0 && money%arr[0] == 0 {
		m = money / arr[0]
		return m
	}

	temp := money - arr[0]*m
	for i := 1; i < len(arr); i++ {
		if temp < arr[i] {
			continue
		}
		
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
