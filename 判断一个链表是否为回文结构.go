package main

import "fmt"

type cycleNode3 struct {
	Val  int
	Next *cycleNode3
}

func hasCycle1(head *cycleNode3) bool {
	if head == nil {
		fmt.Println("---空链表--")
	}
	curl := head

	has := make(map[int]interface{}, 0)

	if curl != nil && curl.Next != nil {

		if _, ok := has[curl.Val]; !ok {
			has[curl.Val] = ""
		} else {
			fmt.Println("-----存在回文链表----")
		}

	}

	return false

}

func main() {
	head := &cycleNode3{}
	head.Val = 1

	cycleLn1 := &cycleNode3{}
	cycleLn1.Val = 3

	cycleLn2 := &cycleNode3{}
	cycleLn2.Val = 5

	cycleLn3 := &cycleNode3{}
	cycleLn3.Val = 3

	head.Next = cycleLn1
	cycleLn1.Next = cycleLn2
	cycleLn2.Next = cycleLn3

	hasCycle1(head)
}
