package main

import "fmt"

type endDeleteNode struct {
	Val  int
	Next *endDeleteNode
}

func FindKDeletethToTail(head *endDeleteNode, k int) *endDeleteNode {
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
	slow.Next = slow.Next.Next
	return slow

}

// 参考https://zhuanlan.zhihu.com/p/364660395
func main() {
	head := &endDeleteNode{}
	head.Val = 1

	cycleLn1 := &endDeleteNode{}
	cycleLn1.Val = 3

	cycleLn2 := &endDeleteNode{}
	cycleLn2.Val = 5

	cycleLn3 := &endDeleteNode{}
	cycleLn3.Val = 3

	head.Next = cycleLn1
	cycleLn1.Next = cycleLn2
	cycleLn2.Next = cycleLn3
	FindKDeletethToTail(head, 3)

}
