package problems

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	var left, right *TreeNode
	left, right = nil, nil
	if root.Left != nil {
		left = invertTree(root.Left)
	}
	if root.Right != nil {
		right = invertTree(root.Right)
	}

	root.Left, root.Right = right, left

	return root

}
