package main

import "math"

func lengthOfLastWord(s string) int {
	saved := math.MaxInt64
	for i := len(s) - 1; i >= 0; i-- {
		if string(s[i]) != " " {
			saved = i
			break
		}
	}
	word := ""
	for i := saved; i >= 0; i-- {
		if string(s[i]) == " " {
			break
		}
		word += string(s[i])
	}
	return len(word)
}
