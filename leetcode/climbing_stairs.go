package leetcode

// 70. Climbing Stairs
func ClimbStairs(n int) int {
	if n == 1 {
		return 1
	}

	result := 0
	// This is how we do the first 2 steps.
	// Essentially, there is a combo of either:
	// 1 step, 1 step OR 2 step
	mem := []int{1, 1}

	// We can iterate to range n - 1 since we already know the first 2 steps
	for range n - 1 {
		// Add the previous steps
		result = mem[0] + mem[1]
		mem[0] = mem[1]
		mem[1] = result
	}

	return result
}
