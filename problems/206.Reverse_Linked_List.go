package problems

// ReverseList problem 206
func ReverseList(head *ListNode) *ListNode {
	return reverseList(head)
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
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
