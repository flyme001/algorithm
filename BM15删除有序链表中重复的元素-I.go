package main

type deleteNode struct {
	Val  int
	Next *deleteNode
}

func deleteNodeList(head1 *deleteNode) *deleteNode {
	if head1 == nil {
		return nil
	}

	curl := head1
	for curl != nil {
		if curl.Val == curl.Next.Val {
			curl.Next = curl.Next.Next
		} else {
			curl = curl.Next
		}
	}

	return curl
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
