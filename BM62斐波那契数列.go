package main

import "fmt"

// fib(1)=1,fib(2)=1,fib(3)=fib(3-1)+fib(3-2)=2,fib(4)=fib(4-1)+fib(4-2)=3
func Fibonacci(n int) int {
	if n == 0 {
		return 1
	}
	if n == 1 {
		return 1
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}

func main() {
	nums := []int{1, 1, 3, 5}
	for i := 0; i < len(nums); i++ {
		fmt.Println(Fibonacci(i))
	}

	fmt.Println(Fibonacci(6))
}
