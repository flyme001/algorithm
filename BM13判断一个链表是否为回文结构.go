package main

import "fmt"

type cycleNode3 struct {
	Val  int
	Next *cycleNode3
}

// 两种方式第一种是数组法 两外一种双指针
func hasCycle1(head *cycleNode3) bool {
	if head == nil {
		fmt.Println("---空链表--")
	}
	fast := head
	slow := head

	if fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}

	// 奇数无法中转 12321 此时慢指针在3 快指针在1 这里如果此时做反转是 123 > 12 所以后进位
	if fast != nil {
		slow = slow.Next
	}

	curl := slow
	pre := &cycleNode3{}
	for curl != nil {
		temp := curl.Next
		curl.Next = pre
		pre = curl
		curl = temp
	}

	for pre != nil {
		if pre.Val != head.Val {
			return false
		} else {
			pre = pre.Next
			head = head.Next
		}
	}

	return true

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
