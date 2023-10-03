package main

import "testing"

func TestAddBinary(t *testing.T) {
	t.Run("11, 1", func(t *testing.T) {
		s1, s2 := "11", "1"
		got := addBinary(s1, s2)
		want := "100"

		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	})
	t.Run("1, 1", func(t *testing.T) {
		s1, s2 := "1", "1"
		got := addBinary(s1, s2)
		want := "10"

		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	})
	t.Run("1010, 1011", func(t *testing.T) {
		s1, s2 := "1010", "1011"
		got := addBinary(s1, s2)
		want := "10101"

		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("HUGE", func(t *testing.T) {

		//24847893154024981730169397005
		s1 := "10100000100100110110010000010101111011011001101110111111111101000000101111001110001111100001101"
		//526700554598729746900966573811
		s2 := "110101001011101110001111100110001010100001101011101010000011011011001011101111001100000011011110011"
		got := addBinary(s1, s2)
		//551548447752754728631135970816
		want := "110111101100010011000101110110100000011101000101011001000011011000001100011110011010010011000000000"

		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}
