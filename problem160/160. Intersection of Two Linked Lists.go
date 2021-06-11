package problem160

import (
	common "github.com/Brian-Ding/sogleet/common"
)

func Judge() {

}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getIntersectionNode(headA, headB *common.ListNode) *common.ListNode {
	if headA == nil || headB == nil {
		return nil
	}

	countA := 0
	countB := 0

	temp := headA
	for temp != nil {
		countA++
		temp = temp.Next
	}

	temp = headB
	for temp != nil {
		countB++
		temp = temp.Next
	}

	tempA := headA
	tempB := headB

	if countA >= countB {
		delta := countA - countB
		for i := 0; i < delta; i++ {
			tempA = tempA.Next
		}
	} else {
		delta := countB - countA
		for i := 0; i < delta; i++ {
			tempB = tempB.Next
		}
	}

	for tempA != nil {
		if tempA == tempB {
			return tempA
		}

		tempA = tempA.Next
		tempB = tempB.Next
	}

	return nil
}
