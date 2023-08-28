package main

import "testing"

// Example 1:

// Input: s = "Hello World"
// Output: 5
// Explanation: The last word is "World" with length 5.
// Example 2:

// Input: s = "   fly me   to   the moon  "
// Output: 4
// Explanation: The last word is "moon" with length 4.
// Example 3:

// Input: s = "Luffy is still JoyBoy"
// Output: 6
// Explanation: The last word is "JoyBoy" with length 6.

func TestLengthOfLastWord(t *testing.T) {
	t.Run("Hello World ", func(t *testing.T) {
		s := "Hello World "
		got := lengthOfLastWord(s)
		want := 5

		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}

	})

	t.Run("   fly me   to   the moon  ", func(t *testing.T) {
		s := "   fly me   to   the moon  "
		got := lengthOfLastWord(s)
		want := 4

		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}

	})

	t.Run("Luffy is still JoyBoy", func(t *testing.T) {
		s := "Luffy is still JoyBoy"
		got := lengthOfLastWord(s)
		want := 6

		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}

	})
}
