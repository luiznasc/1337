package main

import (
	"fmt"
	"reflect"
	"testing"
)

// Input: list1 = [1,2,4], list2 = [1,3,4]
// Output: [1,1,2,3,4,4]
// Example 2:

// Input: list1 = [], list2 = []
// Output: []
// Example 3:

// Input: list1 = [], list2 = [0]
// Output: [0]
func TestMergeTwoLists(t *testing.T) {
	l1 := &ListNode{1, &ListNode{2, &ListNode{4, nil}}}
	l2 := &ListNode{1, &ListNode{3, &ListNode{4, nil}}}
	got := mergeTwoLists(l1, l2)
	want := &ListNode{1, &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{4, nil}}}}}}

	cp := got
	for {
		fmt.Printf("%v", cp)
		if cp == nil {
			break
		}
		cp = cp.Next
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}
