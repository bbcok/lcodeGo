package dataStruct

// KMP 算法
func KMP(s, p string) int {
	next := GetNext(p)
	m, n, x, y := len(s), len(p), 0, 0
	for x < m && y < n {
		if s[x] == p[y] {
			x++
			y++
		} else if y == 0 {
			x++
		} else {
			y = next[y]
		}
	}
	if y == n {
		return x - y
	} else {
		return -1
	}
}
func GetNext(p string) []int {
	n := len(p)
	if n == 1 {
		return []int{-1}
	}
	next := make([]int, n)
	next[0] = -1
	next[1] = 0
	i, cn := 2, 0
	for i < n {
		if p[i-1] == p[cn] {
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
