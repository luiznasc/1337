package main

import (
	"reflect"
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {
	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	nWant := []int{0, 1, 2, 3, 4}
	want := 5
	got := removeDuplicates(nums)

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
	if !reflect.DeepEqual(nums[:len(nWant)], nWant) {
		t.Errorf("got %v, want %v", nums[:len(nWant)], nWant)
	}
}
