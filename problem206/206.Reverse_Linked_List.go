package problems

import "github.com/Brian-Ding/sogleet/common"

// ReverseList problem 206
func ReverseList(head *common.ListNode) *common.ListNode {
	return reverseList(head)
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *common.ListNode) *common.ListNode {
	if head == nil {
		return nil
	}

	last := head
	next := head.Next

	for next != nil {
		// a --> b --> c
		// b --> a
		next.Next, last, next = last, next, next.Next
	}

	return last
}
