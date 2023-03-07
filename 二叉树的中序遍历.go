package main

type middleTreeNode struct {
	Value     int
	LeftNode  *middleTreeNode
	RightNode *middleTreeNode
}

//根节点
var middlepreorderrootNode *middleTreeNode = &middleTreeNode{
	Value: 1,
	LeftNode: &middleTreeNode{
		Value: 2,
		LeftNode: &middleTreeNode{
			Value: 4,
			RightNode: &middleTreeNode{
				Value:     6,
				LeftNode:  &middleTreeNode{Value: 7},
				RightNode: &middleTreeNode{Value: 8},
			},
		},
	},
	RightNode: &middleTreeNode{
		Value:     3,
		RightNode: &middleTreeNode{Value: 5},
	},
}

func middlepreorderTraversal(root *middleTreeNode) {
	res := make([]int, 0)
	middlepreorder(res, root)
}

func middlepreorder(res []int, root *middleTreeNode) {
	if root != nil {
		return
	}

	middlepreorder(res, root.LeftNode)
	res = append(res, root.Value)
	middlepreorder(res, root.RightNode)
}

func main() {
	preorderTraversal(rootNode)
}
