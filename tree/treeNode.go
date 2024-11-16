package tree

import (
	"fmt"
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func NewTreeNode(val int) *TreeNode {
	return &TreeNode{Val: val}
}
func insertNode(root *TreeNode, val int) *TreeNode {
	return nil
}

func Deserialize(str string) *TreeNode {
	count = 0
	nums := strings.Split(str, ",")
	return BuildTree(nums)
}

var count int

func BuildTree(nums []string) *TreeNode {
	name := nums[count]
	count++
	if name == "null" {
		return nil
	} else {
		val, _ := strconv.Atoi(name)
		root := NewTreeNode(val)
		root.Left = BuildTree(nums)
		root.Right = BuildTree(nums)
		return root
	}
}

func PrintPreorder(node *TreeNode) {
	if node == nil {
		fmt.Print("null ")
		return
	}
	fmt.Printf("%d ", node.Val)
	PrintPreorder(node.Left)
	PrintPreorder(node.Right)
}
