package main

func mySqrt(x int) int {
	count := x
	upperEdge := x
	for {
		if count*count > x {
			upperEdge = count
			count = count / 2
		} else {
			if count*count == x {
				return count
			}
			for i := count; i <= upperEdge; i++ {
				if i*i > x {
					return i - 1
				} else if i*i == x {
					return i
				}
			}
		}
	}
}
