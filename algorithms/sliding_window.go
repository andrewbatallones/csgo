package algorithms

import (
	"fmt"

	"github.com/andrewbatallones/csgo/leetcode"
)

// Sliding Window Algorithm
// An example is finding the max of sub array m in array n
// An approach would be to get the sum of each sub array.
// Sliding door: The only real change is the first/last element in the sub array
// Notes to remember:
// - Start is usually going to be either -1 or at the index.
// - End is usually going to be the position, so add a +1.
func SlidingWindow() {
	fmt.Printf("%d\n", exampleMaxSubArr([]int{8, 3, -2, 4, 5, -1, 0, 5, 3, 9, -6}, 5))
	leetcode.PrintSlice(exampleLongestSubArray([]int{4, 5, 2, 0, 1, 8, 12, 3, 6, 9}, 15))
}

// Given an array and the sub array length, find the max sub array
func exampleMaxSubArr(nums []int, subLength int) int {
	maxSum := 0
	val := 0
	start := 0
	end := min(subLength, len(nums))

	// Get init sum
	for i := 0; i < end; i++ {
		maxSum += nums[i]
		// We need to set the initial place
		val = maxSum
	}

	for end < len(nums) {
		// get new value
		// We are still at the start and the end is the len, not the position
		// We can do this since the start is what's being subtracted and the incoming end
		// is what we are adding.
		val = val - nums[start] + nums[end]

		// calculate new max
		maxSum = max(maxSum, val)

		// move slider
		start++
		end++
	}

	return maxSum
}

// In this example, we want to find the largest subarray that would
// be close to the largest provided number (without going over)
func exampleLongestSubArray(nums []int, largest int) []int {
	if len(nums) == 0 {
		return nums
	}

	// Initial starting position
	start := 0
	end := 1
	val := nums[start]

	// initial max vals
	maxVal := min(val, largest) // We get the min in case the init val is larger than largest
	maxStart := 0
	maxEnd := 0

	for end < len(nums) {
		// Check if sub array is smaller than the largest num
		if val <= largest {
			if val > maxVal {
				maxStart = start
				maxEnd = end
				maxVal = val
			}

			end++
			// Add to value
			if end < len(nums) {
				val += nums[end]
			}
		} else {
			val -= nums[start]
			start++

			if start >= end {
				end++
				val += nums[end]
			}
		}
	}

	// Want to add +1 to both.
	// start: technically we started at -1
	// end: it is the index, so we want to return the end length
	return nums[maxStart+1 : maxEnd+1]
}
