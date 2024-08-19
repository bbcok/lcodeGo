package windowSlide

func LengthOfLongestSubstring(s string) int {
	last := make(map[byte]int)
	ans := 0
	for l, r := 0, 0; r < len(s); r++ {
		l = max(l, last[s[r]]+1)
		ans = max(ans, r-l+1)
		last[s[r]] = r
	}
	return ans
}

// func LengthOfLongestSubstring1(s string) int {
// 	last := make(map[byte]int)
// 	ans := 0
// 	for l, r := 0, 0; r < len(s); r++ {

// 	}
// 	return ans
// }
