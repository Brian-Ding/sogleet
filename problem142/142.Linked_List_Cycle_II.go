package problems

import "github.com/Brian-Ding/sogleet/common"

// DetectCycle problem 142
// not pass
func DetectCycle(head *common.ListNode) *common.ListNode {
	return detectCycle(head)
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func detectCycle(head *common.ListNode) *common.ListNode {
	fast := head
	slow := head

	for fast != nil && fast.Next != nil {
		if fast.Next == slow.Next {
			return fast.Next
		}

		fast = fast.Next.Next
		slow = slow.Next
	}

	return nil
}
