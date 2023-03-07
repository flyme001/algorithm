package main

func jumpFloor(n int) int {
	if n == 0 {
		return 1
	}
	if n == 1 {
		return 1
	}
	return jumpFloor(n-1) + jumpFloor(n-2)
}

func main() {
	nums := []int{1, 1, 3, 5}
	for i := 0; i < len(nums); i++ {
		jumpFloor(i)
	}
}
