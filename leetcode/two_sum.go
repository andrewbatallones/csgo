package leetcode

// 1. Two Sum
func TwoSum(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}

	return []int{0, 0}
}

// 167. Two Sum II - Input Array Is sorted
// There will always be at least 2 numbers
func TwoSumTwo(numbers []int, target int) []int {
	// We have 2 pointers (left, right)
	// If the number is bigger, lower right
	// If the number is smaller, increase left

	// Return left,right
	left := 0
	right := len(numbers) - 1

	for left < right {
		check := numbers[left] + numbers[right]

		if check == target {
			break
		} else if check > target {
			right--
		} else {
			left++
		}
	}

	return []int{left + 1, right + 1}
}
