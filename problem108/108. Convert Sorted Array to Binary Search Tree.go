package problem108

import (
	"fmt"

	"github.com/Brian-Ding/sogleet/common"
)

func Judge() {
	array := []int{-10, -3, 0, 5, 9}
	result := sortedArrayToBST(array)

	fmt.Println(result)
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sortedArrayToBST(nums []int) *common.TreeNode {
	if len(nums) == 0 {
		return nil
	}

	if len(nums) == 1 {
		return &common.TreeNode{Val: nums[0]}
	}

	root := divideAndConque(nums)
	return root
}

func divideAndConque(nums []int) *common.TreeNode {
	if len(nums) == 0 {
		return nil
	}

	center := int(len(nums) / 2)

	node := &common.TreeNode{Val: nums[center]}
	node.Left = divideAndConque(nums[:center])
	node.Right = divideAndConque(nums[center+1:])

	return node
}
