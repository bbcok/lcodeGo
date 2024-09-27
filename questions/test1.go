package main

import (
	"fmt"
	"lcodego/dataStruct"
)

func LinkedListInBinaryTree(head *dataStruct.ListNode, root *dataStruct.TreeNode) bool {
	var list []int = make([]int, 0)
	for head != nil {
		list = append(list, head.Val)
		head = head.Next
	}
	fmt.Println(list)
	next := getNextint(list, len(list))
	return find(list, next, root, 0)
}
func find(s2 []int, next []int, root *dataStruct.TreeNode, i int) bool {
	if root == nil {
		return false
	}
	if i == len(s2) {
		return true
	}
	for i >= 0 && root.Val != s2[i] {
		i = next[i]
	}
	return find(s2, next, root.Left, i+1) || find(s2, next, root.Right, i+1)
}
func getNextint(ints []int, index int) []int {
	next := make([]int, len(ints))
	if index == 1 {
		return []int{-1}
	}
	next[0] = -1
	next[1] = 0
	i, cn := 2, 0
	for i < index {
		if ints[i-1] == ints[cn] {
			cn++
			next[i] = cn
			i++
		} else if cn > 0 {
			cn = next[cn]
		} else {
			next[i] = 0
			i++
		}
	}
	return next

}

func DeleteAgainAndAgain(s string, p string) {
	stack1 := make([]int, 0)
	stack2 := make([]int, 0)

	next := dataStruct.GetNext(p)
	n, m, x, y := len(s), len(p), 0, 0
	size := 0
	for x < n {
		if s[x] == p[y] {
			stack1 = append(stack1, x)
			stack2 = append(stack2, y)
			size++
			x++
			y++
		} else if y == 0 {
			stack1 = append(stack1, x)
			stack2 = append(stack2, -1)
			size++
			x++
		} else {
			y = next[y]
		}
		if y == m {
			size = size - m
			if size < 0 {
				y = 0
			} else {
				y = stack2[size-1] + 1
			}

		}
	}

}
func main() {

	//fmt.Println(evalRPN([]string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"}))
	fmt.Println(maxSlidingWindow([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3))
}
