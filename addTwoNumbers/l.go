package main

/**
 * Definition for singly-linked list.
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func twoNumbersRecursive(l1, l2 *ListNode, carry int) *ListNode {
	sum, nextCarry := sumWithCarry(l1.Val, l2.Val)
	if l1.Next == nil && l2.Next == nil {
		if carry != 0 && sum+carry >= 10 {
			sum, carry = sumWithCarry(sum, carry)
			return &ListNode{sum, &ListNode{carry + nextCarry, nil}}
		} else {
			return &ListNode{sum + carry, nil}
		}
	}
	if l2.Next == nil {
		if carry != 0 && sum+carry >= 10 {
			sum, carry = sumWithCarry(sum, carry)
			return &ListNode{sum, twoNumbersRecursive(l1.Next, &ListNode{0, nil}, carry+nextCarry)}
		} else {
			return &ListNode{l1.Val + carry, l1.Next}
		}
	}
	if l1.Next == nil {
		if carry != 0 && sum+carry >= 10 {
			sum, carry = sumWithCarry(sum, carry)
			return &ListNode{sum, twoNumbersRecursive(&ListNode{0, nil}, l2.Next, carry+nextCarry)}
		} else {
			return &ListNode{l2.Val + carry, l2.Next}
		}
	}
	if sum+carry >= 10 {
		sum, carry = sumWithCarry(sum, carry)
		return &ListNode{sum, twoNumbersRecursive(l1.Next, l2.Next, carry+nextCarry)}
	}

	return &ListNode{sum + carry, twoNumbersRecursive(l1.Next, l2.Next, nextCarry)}
}

func sumWithCarry(n1, n2 int) (sum, carry int) {
	sum = n1 + n2
	carry = sum / 10
	sum = sum % 10

	return sum, carry
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	return twoNumbersRecursive(l1, l2, 0)
}
