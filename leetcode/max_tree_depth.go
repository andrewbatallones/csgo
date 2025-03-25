package leetcode

// Given a binary tree, get the max depth
// Going to be a recursive to get the left/right.
func MaxDepth(root *TreeNode) int {
	// If the root is nil => 0
	// If the children are nil => 1
	// else => get left/right recursively (dfs)
	// 		- Check to see which side goes deeper and add it to the side
	if root == nil {
		return 0
	} else if root.Left == nil && root.Right == nil {
		return 1
	} else {
		left := MaxDepth(root.Left)
		right := MaxDepth(root.Right)

		if left > right {
			return left + 1
		} else {
			return right + 1
		}
	}
}
