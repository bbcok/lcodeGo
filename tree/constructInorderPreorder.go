package tree

import "slices"

func constructInorderPreorder(pre []int, in []int) *TreeNode {
	if len(pre) == 0 {
		return nil
	}
	val := pre[0]
	leftSize := slices.Index(in, val)
	root := TreeNode{Val: pre[0]}
	root.Left = constructInorderPreorder(pre[1:leftSize+1], in[:leftSize])
	root.Right = constructInorderPreorder(pre[leftSize+1:], in[leftSize+1:])
	return &root
}
