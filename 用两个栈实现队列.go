package main

var stack1, stack2 []int

func pop(node []int) {
	for i := 0; i < len(node); i++ {
		stack1 = append(stack1, node[i])
	}
}

func push(node1 []int) []int {
	for i := 0; i < len(node1); i++ {
		stack2 = append(stack2, node1[len(node1)-1-i])
	}
	return stack2
}

func main() {

	node := []int{2, 3, 4, 6}
	pop(node)
	push(stack1)

}
