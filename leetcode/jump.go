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
	fmt.Printf("%v\n", CanJump([]int{2, 3, 1, 1, 4}))
	fmt.Printf("%v\n", CanJump([]int{4}))
	fmt.Printf("%v\n", CanJump([]int{3, 2, 1, 0, 4}))
	fmt.Printf("%v\n", CanJump([]int{0, 0, 0, 0, 100}))
	fmt.Printf("%v\n", CanJump([]int{1000, 0, 0, 0, 0}))
}

// 45. Jump Game II
// In this version, we want the min amount of jumps it takes to get to the end
// All nums are able to be completed
// Solution: Greedy BFS
func JumpII(nums []int) int {
	end := len(nums) - 1
	// Current index
	i := 0
	maxReachable := 0
	lastJumpPos := 0
	jumps := 0

	for lastJumpPos < end {
		// Get max how far we can jump
		maxReachable = max(maxReachable, i+nums[i])

		// If the current index == lastJumpPos
		if i == lastJumpPos {
			lastJumpPos = maxReachable
			jumps++
		}

		i++
	}

	return jumps
}

func TestJumpII() {
	fmt.Printf("%v\n", JumpII([]int{2, 3, 1, 1, 4}))
	fmt.Printf("%v\n", JumpII([]int{4}))
	fmt.Printf("%v\n", JumpII([]int{2, 3, 0, 1, 4}))
	fmt.Printf("%v\n", JumpII([]int{1000, 0, 0, 0, 0}))
}
