package tree

import "container/list"

func findBottomLeftValue(root *TreeNode) int {
	queue := list.New()
	queue.PushBack(root)
	leftNode := 0
	for queue.Len() > 0 {
		len1 := queue.Len()
		for i := 0; i < len1; i++ {
			node := queue.Front().Value.(*TreeNode)
			if i == 0 {
				leftNode = node.Val
			}
			queue.Remove(queue.Front())
			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
		}
	}
	return leftNode
}
