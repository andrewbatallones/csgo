package leetcode

// 238. Product of Array Except Self
// Needs to be O(n)
// Cannot divide
func ProductExceptSelf(nums []int) []int {
	// Essentially, we are going to calculate all the products going forward, then backwards.
	// Doing it forwards first will get us the initial products
	// - getting the product of all numbers prior to position.
	// - we can keep track of that in a variable
	// Doing it backwards by multiplying the postFix and the calculated prefix.
	// - result[len(nums)-1] => just the prefix since the initial val of the postFix will be 1
	// - then we need to increase it by the idx of the nums slice
	// Doing that, that should get us the produc and doesn't hit the current index going forward, then backwards.
	result := make([]int, len(nums))

	// Calculate all prefixes (we don't need to get the last one)
	// This is to first calculate the prefix products
	preFix := 1
	for idx, num := range nums {
		result[idx] = preFix
		preFix *= num
	}

	// We need to go from the end to the beginning now
	// This is to calculate the postfix products
	postFix := 1
	for reverseIdx := len(nums) - 1; reverseIdx >= 0; reverseIdx-- {
		result[reverseIdx] = postFix * result[reverseIdx]
		postFix *= nums[reverseIdx]
	}

	return result
}
