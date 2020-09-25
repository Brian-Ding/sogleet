package problems

func SwapPairs(head *ListNode) *ListNode {
	return swapPairs(head)
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	if head.Next == nil {
		return head
	}

	newHead := head.Next

	var last *ListNode
	next := head.Next

	count := 0

	for next != nil {
		if count%2 == 0 {
			if last != nil {
				last.Next = next
			}
			next.Next, head.Next = head, next.Next
			last, next = next, head.Next
		} else {
			last, head, next = head, next, next.Next
		}

		count++
	}

	return newHead
}
