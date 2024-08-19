package main

import (
	"fmt"
	"lcodego/dataStruct"
	"lcodego/windowSlide"
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
	fmt.Println(windowSlide.MinimumWindowSubstring("ADOBEABCODEBANC", "ABC"))
}
