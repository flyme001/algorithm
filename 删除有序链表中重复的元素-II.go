package main

type deleteTwoNode struct {
	Val  int
	Next *deleteTwoNode
}

func deleteTwoNodeList(head1 *deleteTwoNode) *deleteTwoNode {
	if head1 == nil {
		return nil
	}

	slow, fast := head1, head1.Next
	for fast != nil {
		if slow.Val == fast.Val {

			x := slow.Val
			for slow.Next != nil && slow.Next.Val == x {
				slow.Next = slow.Next.Next
			}

		}

		slow = slow.Next
		fast = fast.Next

	}

	return slow
}

func main() {
	head1 := &deleteTwoNode{}
	head1.Val = 3

	ln1 := &deleteTwoNode{}
	ln1.Val = 4

	ln2 := &deleteTwoNode{}
	ln2.Val = 10

	ln3 := &deleteTwoNode{}
	ln3.Val = 10

	head1.Next = ln1
	ln1.Next = ln2

	deleteTwoNodeList(head1)

}
