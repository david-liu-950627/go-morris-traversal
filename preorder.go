/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func preorderTraversal(root *TreeNode) []int {
	current := root
	result := []int{}

	for current != nil {
		result = append(result, current.Val)

		if current.Left == nil {
			current = current.Right
		} else {
			predecessor := current.Left
			for predecessor.Right != nil && predecessor.Right != current {
				predecessor = predecessor.Right
			}

			if predecessor.Right == nil {
				predecessor.Right = current.Right
				current = current.Left
			} else {
				predecessor.Right = nil
				current = current.Right
			}
		}
	}

	return result
}
