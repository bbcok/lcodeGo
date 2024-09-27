package main

func removeDuplicates(s string) string {
	l := make([]byte, 0)
	l = append(l, s[0])
	for i := 1; i < len(s); i++ {
		if len(l) > 0 && l[len(l)-1] == s[i] {
			l = l[:len(l)-1]
		} else {
			l = append(l, s[i])
		}
	}
	return string(l)
}
