package main

import "testing"

// Example 1:

// Input: nums = [1,3,5,6], target = 5
// Output: 2
// Example 2:

// Input: nums = [1,3,5,6], target = 2
// Output: 1
// Example 3:

// Input: nums = [1,3,5,6], target = 7
// Output: 4
func TestSearchInsert(t *testing.T) {
	t.Run("[1,3,5,6] 5", func(t *testing.T) {
		nums := []int{1, 3, 5, 6}

		got := searchInsert(nums, 5)

		want := 2

		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}

	})

	t.Run("[1,3,5,6] 2", func(t *testing.T) {
		nums := []int{1, 3, 5, 6}

		got := searchInsert(nums, 2)

		want := 1

		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}

	})
	t.Run("[1,3,5,6] 7", func(t *testing.T) {
		nums := []int{1, 3, 5, 6}

		got := searchInsert(nums, 7)

		want := 4

		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}

	})
}
