package main

import "testing"

func TestIsValid(t *testing.T) {
	t.Run("(){}}{", func(t *testing.T) {
		s := "(){}}{"
		got := isValid(s)
		want := false

		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	})
	// t.Run("()", func(t *testing.T) {})
}
