package main

import "testing"

func TestMySqrt(t *testing.T) {
	t.Run("4", func(t *testing.T) {
		got := mySqrt(4)
		want := 2
		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	})
	t.Run("8", func(t *testing.T) {
		got := mySqrt(8)
		want := 2
		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}
