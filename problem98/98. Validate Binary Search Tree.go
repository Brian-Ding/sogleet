package problem98

import (
	"fmt"

	"github.com/Brian-Ding/sogleet/common"
)

func Judge() {
	root := &common.TreeNode{Val: 2, Left: &common.TreeNode{Val: 1}, Right: &common.TreeNode{Val: 3}}
	fmt.Println(isValidBST(root))
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isValidBST(root *common.TreeNode) bool {
	if root == nil {
		return false
	}

	array := flat(root)
	if len(array) == 1 {
		return true
	}

	for i := 1; i < len(array); i++ {
		if array[i-1] >= array[i] {
			return false
		}
	}

	return true
}

func flat(node *common.TreeNode) []int {
	array := make([]int, 0)
	if node.Left != nil {
		array = append(array, flat(node.Left)...)
	}

	array = append(array, node.Val)

	if node.Right != nil {
		array = append(array, flat(node.Right)...)
	}

	return array
}
