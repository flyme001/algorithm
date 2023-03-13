package main

import "fmt"

func twoArrSort(arr [][]int) {
	l := len(arr)

	for i := 0; i < l; i++ {
		for j := 0; j < len(arr[i]); j++ {
			fmt.Println(arr[i][j])
		}
	}
}
