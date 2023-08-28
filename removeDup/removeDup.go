package main

func removeDuplicates(nums []int) int {
	currentNumb := -101
	count := 0
	repeated := []int{}
	for _, n := range nums {
		if n != currentNumb {
			repeated = append(repeated, n)
			currentNumb = n
			count++
		}
	}
	copy(nums, repeated)
	// for i, n := range repeated {
	// 	nums[i] = n
	// }
	return count
}
