package tree

import (
	"encoding/json"
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

// PrintPreorder
// 打印先序遍历结果/*
func PrintPreorder(node *TreeNode) {
	if node == nil {
		fmt.Print("null ")
		return
	}
	fmt.Printf("%d ", node.Val)
	PrintPreorder(node.Left)
	PrintPreorder(node.Right)
}

// BuildTreeFromLevelOrder 根据层序遍历结果构造二叉树
func BuildTreeFromLevelOrder(levelOrder []interface{}) *TreeNode {
	if len(levelOrder) == 0 {
		return nil
	}

	// 初始化根节点
	root := &TreeNode{Val: levelOrder[0].(int)}
	queue := []*TreeNode{root}
	i := 1

	for i < len(levelOrder) {
		currentNode := queue[0]
		queue = queue[1:]

		// 处理左子节点
		if levelOrder[i] != nil {
			leftNode := &TreeNode{Val: levelOrder[i].(int)}
			currentNode.Left = leftNode
			queue = append(queue, leftNode)
		}
		i++

		// 处理右子节点
		if i < len(levelOrder) && levelOrder[i] != nil {
			rightNode := &TreeNode{Val: levelOrder[i].(int)}
			currentNode.Right = rightNode
			queue = append(queue, rightNode)
		}
		i++
	}

	return root
}

// BuildTreeLevel 根据层序遍历结果构造二叉树
func BuildTreeLevel(input string) *TreeNode {
	// 将字符串转换为 interface{} 数组
	var levelOrder []interface{}
	err := json.Unmarshal([]byte(input), &levelOrder)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return nil
	}
	for i, v := range levelOrder {
		if v != nil {
			levelOrder[i] = int(v.(float64))
		}
	}

	// 构造二叉树
	root := BuildTreeFromLevelOrder(levelOrder)
	return root
}
