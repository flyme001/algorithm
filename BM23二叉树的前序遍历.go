package main

// 前序遍历：根结点 ---> 左子树 ---> 右子树
//
// 中序遍历：左子树---> 根结点 ---> 右子树
//
// 后序遍历：左子树 ---> 右子树 ---> 根结点
type treeNode struct {
	Value     int
	LeftNode  *treeNode
	RightNode *treeNode
}

// 根节点
var rootNode *treeNode = &treeNode{
	Value: 1,
	LeftNode: &treeNode{
		Value: 2,
		LeftNode: &treeNode{
			Value: 4,
			RightNode: &treeNode{
				Value:     6,
				LeftNode:  &treeNode{Value: 7},
				RightNode: &treeNode{Value: 8},
			},
		},
	},
	RightNode: &treeNode{
		Value:     3,
		RightNode: &treeNode{Value: 5},
	},
}

func preorderTraversal(root *treeNode) {
	res := make([]int, 0)
	preorder(res, root)
}

func preorder(res []int, root *treeNode) {
	if root != nil {
		return
	}
	res = append(res, root.Value)

	preorder(res, root.LeftNode)
	preorder(res, root.RightNode)
}

func main() {
	preorderTraversal(rootNode)
}
