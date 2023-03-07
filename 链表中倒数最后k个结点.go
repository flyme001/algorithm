package main

import "fmt"

type endNode struct {
	Val  int
	Next *endNode
}

func FindKthToTail(head *endNode, k int) *endNode {
	if head == nil {
		fmt.Println("---空链表--")
	}

	fast := head
	slow := head

	for k > 0 {
		fast = fast.Next
		k = k - 1
	}

	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}

	return slow

}

// 参考https://zhuanlan.zhihu.com/p/364660395
func main() {
	head := &endNode{}
	head.Val = 1

	cycleLn1 := &endNode{}
	cycleLn1.Val = 3

	cycleLn2 := &endNode{}
	cycleLn2.Val = 5

	cycleLn3 := &endNode{}
	cycleLn3.Val = 3

	head.Next = cycleLn1
	cycleLn1.Next = cycleLn2
	cycleLn2.Next = cycleLn3
	FindKthToTail(head, 3)

}
