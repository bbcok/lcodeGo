package main

import (
	"fmt"
	"lcodego/dataStruct"
)

func posOrderTraversal(root *dataStruct.TreeNode) {
	list := dataStruct.Stack[dataStruct.TreeNode]{}
	collect := dataStruct.Stack[dataStruct.TreeNode]{}
	list.Push(*root)
	for !list.IsEmpty() {
		head, _ := list.Pop()
		collect.Push(head)
		if head.Left != nil {
			list.Push(*head.Left)
		}
		if head.Right != nil {
			list.Push(*head.Right)
		}
	}
	for !collect.IsEmpty() {
		head, _ := collect.Pop()
		fmt.Println(head.Val)
	}
}
