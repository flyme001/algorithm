package main

var stack []int
var min int

func push43(a int) {
	min = 0
	if a < min {
		min = a
	}
	stack = append(stack, a)
}

func pop43() int {
	res := stack[len(stack)-1]
	stack = stack[:len(stack)-1]
	return res
}

func top43() int {
	if len(stack) != 0 {
		return stack[len(stack)-1]
	}
	return 0
}

func min43() int {
	return min
}

func main() {

}
