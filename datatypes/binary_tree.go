package datatypes

import "fmt"

// A binary tree will have at most 2 nodes.
type BinaryTreeNode struct {
	Val   int
	Left  *BinaryTreeNode
	Right *BinaryTreeNode
}

func NewTree(nums []int) *BinaryTreeNode {
	return insertNode(nums, 0)
}

func insertNode(nums []int, i int) *BinaryTreeNode {
	if i < len(nums) {
		root := &BinaryTreeNode{
			Val: nums[i],
		}

		root.Left = insertNode(nums, 2*i+1)
		root.Right = insertNode(nums, 2*i+2)

		return root
	}

	return nil
}

func (tn *BinaryTreeNode) Print() {
	if tn == nil {
		fmt.Print("")
	} else {
		fmt.Printf("%d ", tn.Val)

		tn.Left.Print()
		tn.Right.Print()
	}
}
