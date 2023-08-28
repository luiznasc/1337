package main

func searchInsert(numbers []int, target int) int {
	p := len(numbers) / 2
	min := 0
	max := len(numbers) - 1
	if numbers[len(numbers)-1] < target {
		return len(numbers)
	}
	for {
		if numbers[p] == target {
			return p
		} else if numbers[p] < target {
			min = p + 1
			p = (min + max) / 2
		} else {
			max = p
			p = (min + max) / 2
		}
		if min == max {
			return p
		}
	}
}
