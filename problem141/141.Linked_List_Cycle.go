package problems

import (
	"math"
)

// HasCycle problem 141
func HasCycle(head *ListNode) bool {
	return hasCycle(head)
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func hasCycle(head *ListNode) bool {
	maxNode := &ListNode{Val: math.MaxInt32}

	current := head
	for current != nil {
		if current.Val == maxNode.Val {
			return true
		}
		current, current.Next = current.Next, maxNode
	}

	return false
}
