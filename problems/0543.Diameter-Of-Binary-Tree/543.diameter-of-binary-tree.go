package problems

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func maxLength(node *TreeNode) (int, int) {
	if node == nil {
		return 0, 0
	}
	mx := 0
	diameter := 0
	leftDepth, leftMax := 0, 0
	rightDepth, rightMax := 0, 0
	if node.Left != nil {
		leftDepth, leftMax = maxLength(node.Left)
		leftDepth++
	}

	if node.Right != nil {
		rightDepth, rightMax = maxLength(node.Right)
		rightDepth++
	}
	mx = max(leftDepth, rightDepth)
	diameter = max(leftDepth+rightDepth, max(leftMax, rightMax))

	return mx, diameter
}

func diameterOfBinaryTree(root *TreeNode) int {
	_, diameter := maxLength(root)
	return diameter
}
