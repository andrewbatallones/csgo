package leetcode

// We are not rotating, but rather moving the elm from the back to the front by k steps
// Example: [1, 2, 3, 4, 5, 6, 7], 3 => [5, 6, 7, 1, 2, 3, 4]
// You can supply k steps larger than the array
func RotateArray(nums []int, k int) {
	// Get the true steps
	step := k % len(nums)
	length := len(nums)

	// we need to shif the array step times
	tmp := make([]int, length)

	for i := 0; i < length; i++ {
		// We are starting at the step and setting the first val
		// We will eventually wrap around the array
		tmp[(i+step)%length] = nums[i]
	}

	copy(nums, tmp)

	PrintSlice(nums)
}
