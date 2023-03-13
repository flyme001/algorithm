package main

import "fmt"

type knode struct {
	Val  int
	Next *knode
}

func merge(head1 *knode, head2 *knode) *knode {
	head := &knode{}
	curl := head

	if head1 == nil {
		curl.Next = head2
	}
	if head2 == nil {
		curl.Next = head1
	}

	for head1 != nil && head2 != nil {
		if head1.Val < head2.Val {
			curl.Next = head1
			head1 = head1.Next
		} else {
			curl.Next = head2
			head2 = head2.Next
		}
		curl = curl.Next
	}
	return curl
}

func kmergeNodeList(list []*knode) {
	var lists *knode
	n := len(list)
	if n == 0 {
		fmt.Println("--nil--")
	}

	if n == 1 {
		fmt.Println(list[0])
	}
	if n == 2 {
		lists = merge(list[0], list[1])
	}

	for i := 2; i < n; i++ {
		lists = merge(lists, list[i])
	}
}

func main() {
	head1 := &knode{}
	head1.Val = 3

	ln1 := &knode{}
	ln1.Val = 4

	ln2 := &knode{}
	ln2.Val = 10

	head2 := &knode{}
	head2.Val = 5

	l3 := &knode{}
	l3.Val = 7

	head1.Next = ln1
	ln1.Next = ln2

	head2.Next = l3
	list := []*knode{}
	list = append(list, head1)
	list = append(list, head2)

	kmergeNodeList(list)

}
