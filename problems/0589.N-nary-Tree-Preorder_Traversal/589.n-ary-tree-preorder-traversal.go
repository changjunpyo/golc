package problems

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func preorder(root *Node) []int {
	if root == nil{
		return []int{}
	}

	l := make([]int,0)
	l = append(l, root.Val)
	for _, node := range root.Children{
		l = append(l, preorder(node)...)
	}
	return l
}
