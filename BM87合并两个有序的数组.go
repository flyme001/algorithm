package main

import "fmt"

func mergeArr() {

	B := []int{1, 2, 3}
	A := make([]int, 6)
	A = append(A, 4, 5, 6)
	m := 3
	n := 3

	for m != 0 && n != 0 {
		if A[m-1] > B[n-1] {
			A[m+n-1] = A[m-1]
			m = m - 1
		} else {
			A[m+n-1] = B[n-1]
			n = n - 1
		}
	}

	fmt.Println("----", A)
	return
}

func main() {
	mergeArr()
}
