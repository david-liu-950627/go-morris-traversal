/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func postorderTraversal(root *TreeNode) []int {
	result := []int{}

	dummyNode := &TreeNode{}
	dummyNode.Left = root
	current := dummyNode

	var previous, successor, temp *TreeNode

	for current != nil {
		if current.Left == nil {
			current = current.Right
		} else {
			predecessor := current.Left
			for predecessor.Right != nil && predecessor.Right != current {
				predecessor = predecessor.Right
			}

			if predecessor.Right == nil {
				predecessor.Right = current
				current = current.Left
			} else {
				predecessor.Right = nil
				successor = current
				current = current.Left
				previous = nil

				for current != nil {
					temp = current.Right
					current.Right = previous
					previous = current
					current = temp
				}

				for previous != nil {
					result = append(result, previous.Val)
					temp = previous.Right
					previous.Right = current
					current = previous
					previous = temp
				}

				current = successor
				current = current.Right
			}
		}
	}

	return result
}
