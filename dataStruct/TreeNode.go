package dataStruct

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func NewTreeNode(val int) *TreeNode {
	return &TreeNode{Val: val}
}
func insertNode(root *TreeNode, val int) *TreeNode {
	return nil
}
