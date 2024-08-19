package windowSlide

func MinimumWindowSubstring(s string, t string) string {
	debetList := make(map[byte]int)
	debet := len(t)
	for _, v := range t {
		debetList[byte(v)] = -1
	}
	ans := int(^uint(0) >> 1)
	start := 0
	for l, r := 0, 0; r < len(s); r++ {
		// if _, ok := debetList[s[r]]; ok {

		if debetList[s[r]] < 0 {
			debetList[s[r]]++
			debet--
		}
		// }
		if debet == 0 {
			for debetList[s[l]] > 0 {
				debetList[s[l]]--
				l++
			}
			if r-l+1 < ans {
				ans = r - l + 1
				start = l
			}
		}
	}
	if ans == int(^uint(0)>>1) {
		return ""
	} else {
		return s[start : start+ans]
	}
}
