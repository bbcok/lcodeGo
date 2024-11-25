package tree

//404. 左叶子之和

func sumOfLeftLeaves(root *TreeNode) int {
	if root == nil {
		return 0
	}

	l := sumOfLeftLeaves(root.Left)
	if root.Left != nil && root.Left.Left == nil && root.Left.Right == nil {
		l += root.Left.Val
	}
	r := sumOfLeftLeaves(root.Right)
	return l + r
}
