package tree

import (
	"container/list"
	"lcodego/dataStruct"
)

func inOrderTraversal(root *dataStruct.TreeNode) []int {
	st := list.New()
	st.PushBack(root)
	res := []int{}
	cur := st.Back().Value.(*dataStruct.TreeNode)
	for root != nil && st.Len() > 0 {

		if cur.Left != nil {
			st.PushBack(cur.Left)
			cur = cur.Left
		} else {
			cur = st.Back().Value.(*dataStruct.TreeNode)
			res = append(res, cur.Val)
			st.Remove(st.Back())
			cur = cur.Right
		}
	}
	return res
}
