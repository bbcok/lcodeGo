package main

import (
	"container/list"
	"fmt"
	"lcodego/dataStruct"
	"lcodego/tree"
)

func testKmp() {
	s := "BBCABCDABABCDABCDABDE"
	p := "ABCDABD"
	re := dataStruct.KMP(s, p)
	fmt.Println(re)

}
func testBinerTree() {
	tree1 := dataStruct.NewTreeNode(1)
	fmt.Print(tree1)
}

func testList() {
	// 创建一个新的双向链表
	l := list.New()

	// 向链表中添加元素
	l.PushBack(1)
	l.PushBack(2)
	l.PushBack(3)

	// 遍历链表并打印元素
	fmt.Println("Initial list:")
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}

	// 插入元素
	frontElement := l.PushFront(0)
	fmt.Println("After inserting 0 at the front:")
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}

	// 删除元素
	l.Remove(frontElement)
	fmt.Println("After removing the first element:")
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}

	// 查找并删除特定元素
	found := false
	for e := l.Front(); e != nil; e = e.Next() {
		if e.Value == 2 {
			l.Remove(e)
			found = true
			break
		}
	}
	if found {
		fmt.Println("After removing 2:")
		for e := l.Front(); e != nil; e = e.Next() {
			fmt.Println(e.Value)
		}
	} else {
		fmt.Println("Element not found.")
	}

	// 获取链表长度
	length := l.Len()
	fmt.Println("Length of the list:", length)
}
func testTree() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	treeNode := dataStruct.ConstructTree(nums)
	res := tree.PreorderTraversal(treeNode)
	str := ""
	for _, v := range res {
		str += fmt.Sprintf("%d ", v)
	}
	fmt.Println(str)
}
func TestFindMode() {
	str := "[0]"
	node := tree.BuildTreeLevel(str)
	rtn := tree.FindMode(node)
	fmt.Println(rtn)
}

func main() {
	// reader := bufio.NewReader(os.Stdin)
	// s := make([]byte, 3)
	// fmt.Fscan(reader, &s)
	// fmt.Println(string(s))
	// fmt.Println(windowSlide.CanCompleteCircuit([]int{1, 2, 3, 4, 5}, []int{3, 4, 5, 1, 2}))
	//s := dataStruct.Stack[string]{}
	//s.Push("1wq2")
	//s.Push("2wq2")
	//s.Pop()
	//fmt.Println(s)
	TestFindMode()
}
