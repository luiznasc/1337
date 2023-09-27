package main

import (
	"reflect"
	"testing"
)

// Example 1:

// Input: digits = [1,2,3]
// Output: [1,2,4]
// Explanation: The array represents the integer 123.
// Incrementing by one gives 123 + 1 = 124.
// Thus, the result should be [1,2,4].
// Example 2:

// Input: digits = [4,3,2,1]
// Output: [4,3,2,2]
// Explanation: The array represents the integer 4321.
// Incrementing by one gives 4321 + 1 = 4322.
// Thus, the result should be [4,3,2,2].
// Example 3:

// Input: digits = [9]
// Output: [1,0]
// Explanation: The array represents the integer 9.
// Incrementing by one gives 9 + 1 = 10.
// Thus, the result should be [1,0].

func TestPlusOne(t *testing.T) {
	t.Run("1, 2, 3", func(t *testing.T) {

		digits := []int{1, 2, 3}
		want := []int{1, 2, 4}

		got := plusOne(digits)

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
	t.Run("4,3,2,1", func(t *testing.T) {

		digits := []int{4, 3, 2, 1}
		want := []int{4, 3, 2, 2}

		got := plusOne(digits)

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
	t.Run("4,3,2,1", func(t *testing.T) {

		digits := []int{9}
		want := []int{1, 0}

		got := plusOne(digits)

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}
