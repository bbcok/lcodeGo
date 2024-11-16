package tree

func countCompleteTree(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := root.Left
	right := root.Right
	leftDepth := 0
	rightDepth := 0
	for left.Left != nil {
		left = left.Left
		leftDepth++
	}
	for right.Right != nil {
		right = right.Right
		rightDepth++
	}
	if leftDepth == rightDepth {
		return (2 << leftDepth) - 1
	}
	return countCompleteTree(root.Right) + countCompleteTree(root.Left) + 1
}
