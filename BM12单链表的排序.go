package main

//https://mojotv.cn/algorithm/golang-quick-sort
type sortNode struct {
	Val  int
	Next *sortNode
}

func sortNodeList(head *sortNode) {
	if head == nil {
		return
	}

	//for h := head; h.Next != nil; h = h.Next {
	//	min := h
	//	for p := h; p != nil; p = p.Next{
	//		if p.Val < min.Val {
	//			min = p
	//		}
	//	}
	//
	//	temp := min.Val
	//	min.Val = h.Val
	//	h.Val = temp
	//
	//}
	//
	for h := head; h.Next != nil; h = h.Next {
		min := h
		for p := h; p != nil; p = p.Next {
			if p.Val < min.Val {
				min = p
			}
		}

		tmp := min.Val
		min.Val = h.Val
		h.Val = tmp
	}
}

func main() {
	head1 := &sortNode{}
	head1.Val = 3

	ln1 := &sortNode{}
	ln1.Val = 4

	ln2 := &sortNode{}
	ln2.Val = 10

	head1.Next = ln1
	ln1.Next = ln2
	sortNodeList(head1)
}
