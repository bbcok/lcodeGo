package main

import (
	"fmt"
	"lcodego/dataStruct"
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
func main() {
	// reader := bufio.NewReader(os.Stdin)
	// s := make([]byte, 3)
	// fmt.Fscan(reader, &s)
	// fmt.Println(string(s))
	// fmt.Println(windowSlide.CanCompleteCircuit([]int{1, 2, 3, 4, 5}, []int{3, 4, 5, 1, 2}))
	s := dataStruct.Stack[string]{}
	s.Push("1wq2")
	s.Push("2wq2")
	s.Pop()
	fmt.Println(s)
}
