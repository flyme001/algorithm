package main

import "fmt"

type nodeList struct {
	Val      int
	NodeNext *nodeList
}

func reversBetweenList(head *nodeList, m int, n int) {
	if head == nil {
		fmt.Println("空链")
	}

	var prev *nodeList
	for i := 1; i < m; i++ {
		prev = prev.NodeNext
	}

	curl := prev.NodeNext
	for i := m; i < n; i++ {
		temp := curl.NodeNext
		curl.NodeNext = prev
		prev = curl
		curl = temp
	}
}

func main() {
	head := &nodeList{}
	head.Val = 1

	ln1 := &nodeList{}
	ln1.Val = 2

	ln2 := &nodeList{}
	ln2.Val = 3

	ln3 := &nodeList{}
	ln3.Val = 4

	head.NodeNext = ln1
	ln1.NodeNext = ln2
	ln2.NodeNext = ln3

	reversBetweenList(head, 2, 3)
}
