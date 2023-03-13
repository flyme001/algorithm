package main

import "fmt"

type cycleNode2 struct {
	Val  int
	Next *cycleNode2
}

func EntryNodeOfLoop(head *cycleNode2) *cycleNode2 {
	if head == nil {
		fmt.Println("---空链表--")
	}
	fast := head
	slow := head

	if fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast.Val == slow.Val {
			tmp := head
			for tmp != slow {
				tmp = tmp.Next
				slow = slow.Next
			}
			return tmp
		}
	}

	return nil

}

func main() {
	head := &cycleNode2{}
	head.Val = 1

	cycleLn1 := &cycleNode2{}
	cycleLn1.Val = 3

	cycleLn2 := &cycleNode2{}
	cycleLn2.Val = 5

	cycleLn3 := &cycleNode2{}
	cycleLn3.Val = 3

	head.Next = cycleLn1
	cycleLn1.Next = cycleLn2
	cycleLn2.Next = cycleLn3

	EntryNodeOfLoop(head)
}
