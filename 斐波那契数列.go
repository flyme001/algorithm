package main

func Fibonacci(n int) int {
	if n == 0 {
		return 1
	}
	if n == 1 {
		return 1
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}

func mian() {
	nums := []int{1, 1, 3, 5}
	for i := 0; i < len(nums); i++ {
		Fibonacci(i)
	}
}
