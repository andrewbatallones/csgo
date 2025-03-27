package datatypes

import "fmt"

// A binary tree will have at most 2 nodes.
type MinHeap struct {
	Val    int
	Parent *MinHeap
	Left   *MinHeap
	Right  *MinHeap
}

func NewMinHeap(nums []int) *MinHeap {
	if len(nums) == 0 {
		return nil
	}

	root := &MinHeap{
		Val: nums[0],
	}

	for i := 1; i < len(nums); i++ {
		root.Insert(nums[i])
	}

	return root
}

func (mh *MinHeap) Insert(num int) {
	// Initial setup, we set the queue with only the root
	queue := []*MinHeap{mh}
	node := &MinHeap{
		Val: num,
	}

	for len(queue) > 0 {
		// pop off queue
		current := queue[0]
		queue = queue[1:]

		if current.Left == nil {
			current.Left = node
			node.Parent = current
			break
		}

		if current.Right == nil {
			current.Right = node
			node.Parent = current
			break
		}

		queue = append(queue, current.Left)
		queue = append(queue, current.Right)
	}

	node.Heapify()
}

func (mh *MinHeap) Heapify() {
	if mh.Parent == nil {
		return
	}

	if mh.Val < mh.Parent.Val {
		// Perform swap
		mh.Val, mh.Parent.Val = mh.Parent.Val, mh.Val
	}

	mh.Parent.Heapify()
}

func (mh *MinHeap) Print() {
	if mh == nil {
		fmt.Print("")
	} else {
		fmt.Printf("%d ", mh.Val)

		mh.Left.Print()
		mh.Right.Print()
	}
}
