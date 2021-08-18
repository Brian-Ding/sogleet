package problem230

import (
	"fmt"

	"github.com/Brian-Ding/sogleet/common"
)

func Judge() {
	root := &common.TreeNode{Val: 3, Left: &common.TreeNode{Val: 1, Right: &common.TreeNode{Val: 2}}, Right: &common.TreeNode{Val: 4}}
	k := 1
	result := kthSmallest(root, k)

	fmt.Println(result) // output: 1
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func kthSmallest(root *common.TreeNode, k int) int {
	_, result := iterate(root, k, 0)

	return result
}

func iterate(node *common.TreeNode, k int, count int) (int, int) {
	if node == nil {
		return count, 0
	}

	count, result := iterate(node.Left, k, count)
	if count == k {
		return count, result
	}

	count++
	if count == k {
		return k, node.Val
	}

	count, result = iterate(node.Right, k, count)
	if count == k {
		return k, result
	}

	return count, 0
}
