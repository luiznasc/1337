package main

/**
 * Definition for singly-linked list.
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoListsRecursive(list1, list2 *ListNode) *ListNode {
	if list1 == nil && list2 == nil {
		return nil
	}
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}
	if list1.Val <= list2.Val {
		list1.Next = mergeTwoListsRecursive(list1.Next, list2)
		return list1
	}
	if list2.Val < list1.Val {
		list2.Next = mergeTwoListsRecursive(list2.Next, list1)
		return list2
	}
	return nil
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {

	return mergeTwoListsRecursive(list1, list2)
}
