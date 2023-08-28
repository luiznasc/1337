package main

import (
	"testing"
)

// strs = ["dog","racecar","car"]
// strs = ["flower","flow","flight"]
func TestLongestCommonPrefix(t *testing.T) {
	t.Run("[dog,racecar,car]", func(t *testing.T) {
		strs := []string{"dog", "racecar", "car"}
		got := longestCommonPrefix(strs)
		want := ""

		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("[flower,flow,flight]", func(t *testing.T) {
		strs := []string{"flower", "flow", "flight"}
		got := longestCommonPrefix(strs)
		want := "fl"

		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}
