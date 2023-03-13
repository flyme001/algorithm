package main

import "fmt"

func tempArr(ar []int, m, n int) []int {
	return append(ar[m-n:m], ar[:m-n]...)
}
func main() {
	ar := []int{
		1, 3, 4, 5,
	}
	fmt.Println(tempArr(ar, 4, 2))
}
