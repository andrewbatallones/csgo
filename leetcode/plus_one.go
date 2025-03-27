package leetcode

// 66. Plus One
func PlusOne(digits []int) []int {
	// Get the index of last digit
	idx := len(digits) - 1

	// Set the remainder
	remainder := 0

	digits[idx], remainder = addPlusOne(digits[idx])
	idx--

	for remainder > 0 && idx >= 0 {
		digits[idx], remainder = addPlusOne(digits[idx])
		idx--
	}

	// Shift the entire array if we still have a remainder
	// One thing I didn't know about slices is that you can append slices by adding `...`
	// http://go.dev/doc/go1.17_spec#Passing_arguments_to_..._parameters
	if remainder > 0 {
		digits = append([]int{remainder}, digits...)
	}

	return digits
}

// Adds one to the value, return the remainder if greater than 9
func addPlusOne(digit int) (int, int) {
	digit++

	if digit >= 10 {
		return 0, 1
	}

	return digit, 0
}
