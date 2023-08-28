package main

func isValid(s string) bool {
	open := [10000]rune{}

	i := 0
	for _, r := range s {
		if i == 0 {
			switch r {
			case '}', ']', ')':
				return false
			}
		}
		switch r {
		case '(', '[', '{':
			open[i] = r
			i++
		case ')':
			if open[i-1] != '(' {
				return false
			}
			open[i] = 0
			i--
		case ']':
			if open[i-1] != '[' {
				return false
			}
			open[i] = 0
			i--
		case '}':
			if open[i-1] != '{' {
				return false
			}
			open[i] = 0
			i--
		}
	}
	return i == 0
}
