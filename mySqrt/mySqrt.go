package main

func mySqrt(x int) int {
	count := 0
	for {
		if count*count > x {
			count--
			break
		}
		count++
	}
	return count
}
