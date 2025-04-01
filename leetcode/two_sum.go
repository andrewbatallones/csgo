package leetcode

// 167. Two Sum II - Input Array Is sorted
// There will always be at least 2 numbers
func TwoSum(numbers []int, target int) []int {
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
