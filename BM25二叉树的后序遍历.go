package main

type followTreeNode struct {
	Value     int
	LeftNode  *followTreeNode
	RightNode *followTreeNode
}

//根节点
var followpreorderrootNode *followTreeNode = &followTreeNode{
	Value: 1,
	LeftNode: &followTreeNode{
		Value: 2,
		LeftNode: &followTreeNode{
			Value: 4,
			RightNode: &followTreeNode{
				Value:     6,
				LeftNode:  &followTreeNode{Value: 7},
				RightNode: &followTreeNode{Value: 8},
			},
		},
	},
	RightNode: &followTreeNode{
		Value:     3,
		RightNode: &followTreeNode{Value: 5},
	},
}

func followpreorderTraversal(root *followTreeNode) {
	res := make([]int, 0)
	followpreorder(res, root)
}

func followpreorder(res []int, root *followTreeNode) {
	if root != nil {
		return
	}

	followpreorder(res, root.LeftNode)
	followpreorder(res, root.RightNode)
	res = append(res, root.Value)
}

func main() {
	followpreorderTraversal(followpreorderrootNode)
}

//二叉树部分后面在实际看看
