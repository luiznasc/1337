package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestAddTwoNumbers(t *testing.T) {
	t.Run("[2,4,3], [5,6,7]", func(t *testing.T) {
		l1 := &ListNode{2, &ListNode{4, &ListNode{3, nil}}}
		l2 := &ListNode{5, &ListNode{6, &ListNode{4, nil}}}
		got := addTwoNumbers(l1, l2)
		want := &ListNode{7, &ListNode{0, &ListNode{8, nil}}}
		copy := got
		for {
			fmt.Print(copy)
			if copy.Next == nil {
				break
			}
			copy = copy.Next
		}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %#v, want %#v", got, want)
		}
	})
	t.Run("l1 = [9,9,9,9,9,9,9], l2 = [9,9,9,9]", func(t *testing.T) {
		l1 := &ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, &ListNode{}}}}}}}}
		l2 := &ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, &ListNode{}}}}}
		got := addTwoNumbers(l1, l2)
		// [8,9,9,9,0,0,0,1]
		want := &ListNode{8, &ListNode{9, &ListNode{9, &ListNode{9, &ListNode{0, &ListNode{0, &ListNode{0, &ListNode{1, nil}}}}}}}}
		copy := got
		for {
			fmt.Print(copy)
			if copy.Next == nil {
				break
			}
			copy = copy.Next
		}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %#v, want %#v", got, want)
		}

	})
}
