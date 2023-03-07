package main

type deleteNode struct {
	Val  int
	Next *deleteNode
}

func deleteNodeList(head1 *deleteNode) *deleteNode {
	if head1 == nil {
		return nil
	}

	slow, fast := head1, head1.Next
	for head1 != nil {
		if slow.Val == fast.Val {
			slow.Next = fast
		}
	}

	return slow
}

func main() {
	head1 := &deleteNode{}
	head1.Val = 3

	ln1 := &deleteNode{}
	ln1.Val = 4

	ln2 := &deleteNode{}
	ln2.Val = 10

	head1.Next = ln1
	ln1.Next = ln2

	deleteNodeList(head1)

}
