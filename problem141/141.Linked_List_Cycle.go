package problems

import (
	"math"

	"github.com/Brian-Ding/sogleet/common"
)

// HasCycle problem 141
func HasCycle(head *common.ListNode) bool {
	return hasCycle(head)
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func hasCycle(head *common.ListNode) bool {
	maxNode := &common.ListNode{Val: math.MaxInt32}

	current := head
	for current != nil {
		if current.Val == maxNode.Val {
			return true
		}
		current, current.Next = current.Next, maxNode
	}

	return false
}
