package main

type node struct {
	Val  int
	Next *node
}

func mergeNodeList(head1 *node, head2 *node) *node {
	if head1 == nil {
		return head2
	}
	if head2 == nil {
		return head1
	}

	var pre *node
	for head1 != nil && head2 != nil {
		if head1.Val < head2.Val {
			pre.Next = head1
			head1 = head1.Next
		} else {
			pre.Next = head2
			head2 = head2.Next
		}
		pre = pre.Next
	}

	//另外一种是出现a链表比b链表大所以就是后面没有比较的链表直接挂到a后面即可
	if head1 != nil {
		pre.Next = head1
	} else {
		pre.Next = head2
	}
	return pre
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
