package leetcode

import "fmt"

// 55. Jump Game
// check if the array can make jumps to the end
// val represent max jump distance
func CanJump(nums []int) bool {
	end := len(nums) - 1
	// This will keep track of where we are jumping
	idx := end - 1

	// Want to walk backwards and check jumps
	// If the jump is valid, we set the new "end point"
	for idx >= 0 {
		if idx+nums[idx] >= end {
			// Set the end to the current idx (basically where it jumped from)
			end = idx
		}

		idx--
	}

	return end == 0
}

func TestCanJump() {
	// fmt.Printf("%v\n", CanJump([]int{2, 3, 1, 1, 4}))
	// fmt.Printf("%v\n", CanJump([]int{4}))
	// fmt.Printf("%v\n", CanJump([]int{3, 2, 1, 0, 4}))
	// fmt.Printf("%v\n", CanJump([]int{0, 0, 0, 0, 100}))
	fmt.Printf("%v\n", CanJump([]int{1000, 0, 0, 0, 0}))
}
