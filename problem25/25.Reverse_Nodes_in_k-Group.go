package problems

import (
	"fmt"

	common "github.com/Brian-Ding/sogleet/common"
)

func Judge() {
	head := &common.ListNode{}
	k := 0
	r := reverseKGroup(head, k)
	fmt.Println(r)
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *common.ListNode
 * }
 */
func reverseKGroup(head *common.ListNode, k int) *common.ListNode {
	if head == nil {
		return nil
	}

	if head.Next == nil {
		return head
	}

	length := 1
	lHead := head
	for {
		lHead = lHead.Next
		if lHead != nil {
			length++
		} else {
			break
		}
	}

	if k > length {
		return head
	}

	last := head
	next := head.Next

	for count := 1; next != nil && count < k; count++ {
		temp := next.Next
		next.Next = last
		head.Next = temp

		last = next
		next = temp
	}

	head.Next = reverseKGroup(next, k)

	return last
}
