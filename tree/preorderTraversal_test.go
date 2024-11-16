package tree

import (
	"container/list"
	"fmt"
	"lcodego/dataStruct"
	"testing"
)

func TestPreorderTraversal(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	treeNode := dataStruct.ConstructTree(nums)
	res := PreorderTraversal(treeNode)
	str := ""
	for _, v := range res {
		str += fmt.Sprintf("%d ", v)
	}
	t.Log(str)
}
func TestList(t *testing.T) {
	// 创建一个新的链表
	l := list.New()

	// 添加元素
	l.PushFront(1)
	l.PushBack(2)
	e := l.PushBack(3)
	l.InsertBefore(4, e)
	l.InsertAfter(5, e)

	// 打印链表
	fmt.Println("初始链表:")
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}

	// 删除元素
	element := l.Front()
	l.Remove(element)

	// 打印删除后的链表
	fmt.Println("删除第一个元素后的链表:")
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}

	// 获取链表长度
	length := l.Len()
	fmt.Println("链表长度:", length)
}
func TestBuildTree(t *testing.T) {
	str := "1,2,null,null,3,4,null,null,5,null,null"
	node := Deserialize(str)
	PrintPreorder(node)
}
