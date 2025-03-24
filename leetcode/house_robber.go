package leetcode

import "slices"

// 2560. House Robber IV
// Return the minimum capability the robber has to steal houses.
// nums: Houses
// k: the minimum number of houses the robber needs to rob
func HouseRobberIV(nums []int, k int) int {
	left := slices.Min(nums)
	right := slices.Max(nums)
	result := 0

	for left <= right {
		mid := (left + right) / 2

		// Essentially, we are checking if the mid point is a potential val.
		if isValid(nums, k, mid) {
			// We are able to rob and we set the result and lower the threshold
			result = mid
			right = mid - 1
		} else {
			// Raise threshold if it's invalid
			left = mid + 1
		}
	}

	return result
}

func isValid(nums []int, k, capability int) bool {
	i := 0
	count := 0

	for i < len(nums) {
		if nums[i] <= capability {
			// Increase by 2 since the Robber can't steal from adjecent houses.
			i += 2
			count++
		} else {
			i++
		}

		if count == k {
			break
		}
	}

	return count == k
}
