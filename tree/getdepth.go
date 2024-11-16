package tree

//最大深度
import (
	"lcodego/dataStruct"
)

func maxDepth(root *dataStruct.TreeNode) int {
	var res int
	res = getDepth(root)
	return res
}
func getDepth(root *dataStruct.TreeNode) int {
	if root == nil {
		return 0
	}

	left := getDepth(root.Left)
	right := getDepth(root.Right)
	return max(left, right) + 1
}
