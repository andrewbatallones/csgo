package leetcode

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func NewTree(nums []int) *TreeNode {
	return insertNode(nums, 0)
}

func insertNode(nums []int, i int) *TreeNode {
	if i < len(nums) {
		root := &TreeNode{
			Val: nums[i],
		}

		root.Left = insertNode(nums, 2*i+1)
		root.Right = insertNode(nums, 2*i+2)

		return root
	}

	return nil
}

func (tn *TreeNode) Print() {
	if tn == nil {
		fmt.Print("")
	} else {
		fmt.Printf("%d ", tn.Val)

		tn.Left.Print()
		tn.Right.Print()
	}
}
