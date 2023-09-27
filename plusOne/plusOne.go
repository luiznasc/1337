package main

func plusOne(digits []int) []int {
	lastI := len(digits) - 1
	digits[lastI] += 1
	currentI := lastI
	for currentI >= 0 && digits[currentI] > 9 {
		digits[currentI] = 0
		currentI--
		if currentI >= 0 {
			digits[currentI]++
		}
	}
	if digits[0] == 0 {
		digits = append([]int{1}, digits...)
	}
	return digits
}
