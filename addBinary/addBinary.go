package main

import "strings"

func addBinary(a string, b string) string {
	smaller, bigger, offset := getSmaller(a, b)
	var resultString strings.Builder
	carry := 0
	nextCarry := 0

	for i := len(smaller) - 1; i >= 0; i-- {
		carry = nextCarry
		nextCarry = 0
		if smaller[i] == '1' && bigger[i+offset] == '1' {
			if carry == 1 {
				resultString.WriteString("1")
			} else {
				resultString.WriteString("0")
			}
			nextCarry = 1
		} else if smaller[i] == '0' && bigger[i+offset] == '0' {
			if carry == 1 {
				resultString.WriteString("1")
			} else {
				resultString.WriteString("0")
			}
		} else {
			if carry == 1 {
				resultString.WriteString("0")
				nextCarry = 1
			} else {
				resultString.WriteString("1")
			}
		}
	}

	for i := offset - 1; i >= 0; i-- {
		carry = nextCarry
		nextCarry = 0
		if carry == 1 {
			if bigger[i] == '1' {
				resultString.WriteString("0")
				nextCarry = 1
			} else {
				if carry == 1 {
					resultString.WriteString("1")
				} else {
					resultString.WriteString("0")
				}
			}
		} else {
			resultString.WriteByte(bigger[i])
		}

	}

	if nextCarry == 1 {
		resultString.WriteString("1")
	}
	var reverse strings.Builder
	result := resultString.String()
	for i := resultString.Len() - 1; i >= 0; i-- {
		reverse.WriteByte(result[i])
	}
	return reverse.String()
}

func getSmaller(a, b string) (string, string, int) {
	if len(a) < len(b) {
		offset := len(b) - len(a)
		return a, b, offset
	} else {
		offset := len(a) - len(b)
		return b, a, offset
	}
}
