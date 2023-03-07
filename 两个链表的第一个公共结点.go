package main

type sameNode struct {
	Val  int
	Next *sameNode
}

func sameNodeList(head1 *sameNode, head2 *sameNode) *sameNode {
	if head1 == nil {
		return nil
	}
	if head2 == nil {
		return nil
	}
	curl1, curl2 := head1, head2
	for head1 != nil && head2 != nil {
		if curl1.Val == curl2.Val {
			return curl1
		} else {
			curl2 = curl2.Next
			curl1 = curl1.Next
		}
	}
	return nil
}

func main() {
	head1 := &sameNode{}
	head1.Val = 3

	ln1 := &sameNode{}
	ln1.Val = 4

	ln2 := &sameNode{}
	ln2.Val = 10

	head2 := &sameNode{}
	head2.Val = 5

	l3 := &sameNode{}
	l3.Val = 4

	head1.Next = ln1
	ln1.Next = ln2

	head2.Next = l3
	sameNodeList(head1, head2)

}
