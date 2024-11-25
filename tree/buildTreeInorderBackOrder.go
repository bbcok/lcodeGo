package tree

import "slices"

func buildTreeInorderBackOrder(back []int, in []int) *TreeNode {
	if len(back) == 0 {
		return nil
	}
	backLen := len(back)
	val := back[backLen-1]
	leftSize := slices.Index(in, val)
	root := &TreeNode{Val: val}
	root.Left = buildTreeInorderBackOrder(back[:leftSize], in[:leftSize])
	root.Right = buildTreeInorderBackOrder(back[leftSize:backLen], in[leftSize+1:])
	return root
}
