package problem101

import (
	"fmt"

	common "github.com/Brian-Ding/sogleet/common"
)

func Judge() {
	root := &common.TreeNode{}
	r := isSymmetric(root)

	fmt.Println(r)
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSymmetric(root *common.TreeNode) bool {
	result := isSym(root.Left, root.Right)

	return result
}

func isSym(node1 *common.TreeNode, node2 *common.TreeNode) bool {
	if node1 == nil && node2 == nil {
		return true
	}

	if (node1 == nil && node2 != nil) || (node1 != nil && node2 == nil) {
		return false
	}

	if node1.Val != node2.Val {
		return false
	}

	result := isSym(node1.Left, node2.Right)
	if !result {
		return result
	}

	result = isSym(node1.Right, node2.Left)

	return result
}
