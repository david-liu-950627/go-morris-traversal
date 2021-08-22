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

		if current.Left == nil {
			result = append(result, current.Val)
			current = current.Right
		} else {
			predecessor := current.Left
			for predecessor.Right != nil && predecessor.Right != current {
				predecessor = predecessor.Right
			}

			if predecessor.Right == nil {
				result = append(result, current.Val)
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
