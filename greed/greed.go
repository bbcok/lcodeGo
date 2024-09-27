package greed

import (
	"sort"
)

func FindContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)
	j := 0
	for i := 0; i < len(s); i++ {
		if j < len(g) && g[j] >= s[i] {
			j++
		}
	}
	return j
}
