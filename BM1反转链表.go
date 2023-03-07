package main

import "fmt"

type listNode struct {
	Val  int
	Next *listNode
}

type listNode1 struct {
	Val  int
	Next *listNode
}

func reversList(head *listNode) *listNode {
	var prev *listNode
	curl := head
	if curl == nil {
		fmt.Println("空链表")
	}

	for curl != nil {
		temp := curl.Next
		curl.Next = prev
		prev = curl
		curl = temp
	}

	return prev
}

func main() {

	head := &listNode{}
	head.Val = 1

	ln2 := &listNode{}
	ln2.Val = 2

	ln3 := &listNode{}
	ln3.Val = 3

	head.Next = ln2
	ln2.Next = ln3

	prev := reversList(head)

	for prev != nil {
		fmt.Println("---val--", prev.Val)
		prev = prev.Next //移动指针
	}
}
