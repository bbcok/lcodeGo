package tree

var res1 []int = make([]int, 0)
var preNode *TreeNode
var count1 int = 0
var maxCount int = 0

func FindMode(root *TreeNode) []int {
	dfs3(root)
	return res1
}
func dfs3(root *TreeNode) {
	if root == nil {
		return
	}
	dfs3(root.Left)
	if preNode == nil {
		count1 = 1
	} else if preNode.Val == root.Val {
		count1++
	} else {
		count1 = 1
	}

	if count1 == maxCount {
		res1 = append(res1, root.Val)
	}
	if count1 > maxCount {
		maxCount = count1
		res1 = make([]int, 0)
		res1 = append(res1, root.Val)
	}
	preNode = root
	dfs3(root.Right)
}
