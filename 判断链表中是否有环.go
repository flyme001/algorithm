package main

import "fmt"

type cycleNode1 struct {
	Val  int
	Next *cycleNode1
}

func hasCycle(head *cycleNode1) bool {
	if head == nil {
		fmt.Println("---空链表--")
	}
	fast := head
	slow := head

	if fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast.Val == slow.Val {
			return true
		}
	}

	return false

}

func main() {
	head := &cycleNode1{}
	head.Val = 1

	cycleLn1 := &cycleNode1{}
	cycleLn1.Val = 3

	cycleLn2 := &cycleNode1{}
	cycleLn2.Val = 5

	cycleLn3 := &cycleNode1{}
	cycleLn3.Val = 3

	head.Next = cycleLn1
	cycleLn1.Next = cycleLn2
	cycleLn2.Next = cycleLn3

	hasCycle(head)
}
