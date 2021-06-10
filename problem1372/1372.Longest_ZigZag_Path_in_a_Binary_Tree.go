package problems

// LongestZigZag problem 1372
func LongestZigZag() int {
	a := &TreeNode{1, nil, &TreeNode{2, &TreeNode{3, nil, nil}, &TreeNode{4, &TreeNode{5, nil, &TreeNode{6, nil, &TreeNode{7, nil, nil}}}, &TreeNode{8, nil, nil}}}}
	return longestZigZag(a)
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func longestZigZag(root *TreeNode) int {
	if root == nil {
		return 0
	}

	left := step(root.Left, true, 0)
	right := step(root.Right, false, 0)
	return max(left, right)
}

func step(node *TreeNode, left bool, value int) int {
	if node == nil {
		return value
	}

	sum := 0
	if left {
		sum = step(node.Right, false, value+1)
		sum = max(sum, step(node.Left, true, 0))
	} else {
		sum = step(node.Left, true, value+1)
		sum = max(sum, step(node.Right, false, 0))
	}

	return sum
}
