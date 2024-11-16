package tree

import "math"

func isBalance(root *TreeNode) bool {
	res := deepth(root)
	return res == -1
}

func deepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	l := deepth(root.Left)
	if l == -1 {
		return -1
	}
	r := deepth(root.Right)
	if r == -1 {
		return -1
	}
	if math.Abs(float64(l-r)) > 1 {
		return -1
	}
	return max(l, r) + 1
}
