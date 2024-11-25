package tree

import (
	"strconv"
)

var res []string

func allRoutes(root *TreeNode) []string {
	res = make([]string, 0)
	dfs(root, "")
	return res
}
func dfs(root *TreeNode, s string) {
	if root.Left == nil && root.Right == nil {
		s = s + strconv.Itoa(root.Val)
		res = append(res, s)
	}

	s = s + strconv.Itoa(root.Val) + "->" //隐含的回溯，下面两个递归返回时，s还是进入递归之前的状态
	if root.Left != nil {
		dfs(root.Left, s)

	}
	if root.Right != nil {
		dfs(root.Right, s)
	}

}
