package tree

//最小深度

import (
	"lcodego/dataStruct"
)

var minDepth int

func getMinDepth(root *dataStruct.TreeNode) {
	if root != nil {
		minDepth = 1
	}
	getDepth(root)
}
func getDepthMin(root *dataStruct.TreeNode) int {
	if root == nil {
		return 0
	}
	leftDepth := getDepthMin(root.Left)
	rightDepth := getDepthMin(root.Right)
	if root.Left != nil && root.Right == nil {
		return leftDepth + 1
	}
	if root.Left == nil && root.Right != nil {
		return rightDepth + 1
	}
	return min(leftDepth, rightDepth) + 1
}
