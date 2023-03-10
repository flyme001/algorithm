package main

import (
	"fmt"
)

func mergeArrs(arr [][]int) [][]int {
	l := len(arr)
	if l == 1 {
		return arr
	}

	data1 := [][]int{}
	temp := []int{}

	for _, v := range arr {
		for _, v1 := range v {
			temp = append(temp, v1)
		}
	}

	for k, v := range temp {
		data := []int{}
		if k == 0 {
			data[0] = v
			continue
		}

		if k == len(data)-2 {
			data[1] = temp[k+2]
			data1 = append(data1, data)

		} else {
			if temp[k] > temp[k+1] && temp[k] < temp[k+2] {
				data[1] = temp[k+2]
				data = []int{}
			}
		}
	}

	return data1
}

func main() {
	arr := [][]int{}
	data := []int{
		1, 2, 3, 4,
	}

	for k, _ := range data {
		fmt.Println(k)
		if k == len(data)-2 {
			if data[k] > data[k+1] && data[k] < data[k+2] {

			}
		}

	}
	arr = append(arr, data)
	fmt.Println(len(arr))
	for _, v := range arr {
		fmt.Println(v)

	}
}
