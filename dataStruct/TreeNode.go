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
func ConstructTree(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	mid := len(nums) / 2
	root := NewTreeNode(nums[mid])
	root.Left = ConstructTree(nums[:mid])
	root.Right = ConstructTree(nums[mid+1:])
	return root
}
