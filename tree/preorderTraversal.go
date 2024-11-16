package tree

import (
	"lcodego/dataStruct"
)

func PreorderTraversal(root *dataStruct.TreeNode) []int {
	stack := dataStruct.NewStack[*dataStruct.TreeNode]()
	var res []int = make([]int, 0)
	stack.Push(root)
	for !stack.IsEmpty() {
		node, _ := stack.Pop()
		if node.Right != nil {
			stack.Push(node.Right)
		}
		if node.Left != nil {
			stack.Push(node.Left)

		}
		res = append(res, node.Val)
	}
	return res
}
