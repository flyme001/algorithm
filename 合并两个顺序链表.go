package main

type node struct {
	Val  int
	Next *node
}

func mergeNodeList(head1 *node, head2 *node) {
	head := &node{}
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
}

func main() {
	head1 := &node{}
	head1.Val = 3

	ln1 := &node{}
	ln1.Val = 4

	ln2 := &node{}
	ln2.Val = 10

	head2 := &node{}
	head2.Val = 5

	l3 := &node{}
	l3.Val = 7

	head1.Next = ln1
	ln1.Next = ln2

	head2.Next = l3

	mergeNodeList(head1, head2)
}
