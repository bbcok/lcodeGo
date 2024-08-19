package main

func RepeatMinimumLength(s string) int {
	next := getNext(s)
	return len(s) - next[len(s)]
}
func getNext(s string) []int {
	next := make([]int, len(s)+1)
	next[0] = -1
	i := 0
	j := -1
	for i <= len(s)-1 {
		if j == -1 || s[i] == s[j] {
			i++
			j++
			next[i] = j
		} else {
			j = next[j]
		}
	}
	return next
}
