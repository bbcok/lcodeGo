package tree

var target int

// 112. 路径总和
func hasPathSum(root *TreeNode, targetSum int) bool {
	target = targetSum
	if root == nil {
		return false
	}
	return dfs1(root, 0)
}
func dfs1(root *TreeNode, sum int) bool {
	if root.Left == nil && root.Right == nil {
		if target == sum+root.Val {
			return true
		} else {
			return false
		}
	}
	val := root.Val
	sum = sum + val
	if root.Left != nil {
		if dfs1(root.Left, sum) {
			return true
		}
	}
	if root.Right != nil {
		if dfs1(root.Right, sum) {
			return true
		}
	}
	sum = sum - val
	return false
}

var pathSumRes [][]int

// 113. 路径总和
func pathSum(root *TreeNode, targetSum int) [][]int {
	if root == nil {
		return pathSumRes
	}
	pathSumRes = make([][]int, 0)
	target = targetSum
	dfs2(root, 0, make([]int, 0))
	return pathSumRes
}
func dfs2(root *TreeNode, sum int, list []int) {
	if root.Left == nil && root.Right == nil {
		if target == sum+root.Val {
			pathCopy := make([]int, len(list))
			for i, element := range list {
				pathCopy[i] = element
			}
			pathCopy = append(pathCopy, root.Val)
			pathSumRes = append(pathSumRes, pathCopy)
		}
		return
	}
	val := root.Val
	sum = sum + val
	list = append(list, root.Val)
	if root.Left != nil {
		dfs2(root.Left, sum, list)
	}
	if root.Right != nil {
		dfs2(root.Right, sum, list)
	}
	sum = sum - val
	list = list[:len(list)-1]

}
