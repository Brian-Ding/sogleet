package problem24

import (
	"fmt"

	common "github.com/Brian-Ding/sogleet/common"
)

func Judge() {
	head := &common.ListNode{}
	r := swapPairs(head)
	fmt.Println(r)
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *common.ListNode) *common.ListNode {
	if head == nil {
		return nil
	}

	if head.Next == nil {
		return head
	}

	newHead := head.Next

	var last *common.ListNode
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
