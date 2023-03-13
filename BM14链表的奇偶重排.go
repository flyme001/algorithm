package main

type oddEvenNode struct {
	Val  int
	Next *oddEvenNode
}

func oddEvenList(head1 *oddEvenNode) *oddEvenNode {
	if head1 == nil {
		return nil
	}
	oddNode := &oddEvenNode{}
	evenNode := &oddEvenNode{}

	temp := &oddEvenNode{
		Val: head1.Val,
	}

	for head1 != nil {
		if head1.Val%2 == 0 {
			oddNode.Next = temp
			oddNode = temp
		} else {
			evenNode.Next = temp
			evenNode = temp
		}
		head1 = head1.Next
	}

	oddNode.Next = evenNode
	return nil
}

func main() {
	head1 := &oddEvenNode{}
	head1.Val = 3

	ln1 := &oddEvenNode{}
	ln1.Val = 4

	ln2 := &oddEvenNode{}
	ln2.Val = 10

	head1.Next = ln1
	ln1.Next = ln2

	oddEvenList(head1)

}
