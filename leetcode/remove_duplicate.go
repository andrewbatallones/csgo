package leetcode

import "fmt"

// 26. Remove Duplicates from Sorted Array
// Takes in an array of nums
// Returns the new size of the array
func RemoveDuplicates(nums []int) int {
	var lastNum int
	count := 0
	left := 0
	right := len(nums) - 1

	for left <= right {
		if right < 0 {
			break
		} else if left == 0 {
			// Set init val, then move next spot
			lastNum = nums[left]
			count++
			left++
		} else if lastNum == nums[left] {
			// move val to end
			// DONT increase count
			// Recheck
			MoveBack(nums, left)
			right--
			// loop again and re-test moved val
		} else {
			// assign new num and go to next val
			// increase count
			lastNum = nums[left]
			left++
			count++
		}

	}

	return count
}

// now each element needs to be there at most twice
func RemoveDuplicatesII(nums []int) int {
	count := 1
	saw := false

	for idx := 1; idx < len(nums); idx++ {
		if nums[idx] == nums[idx-1] && !saw {
			nums[count] = nums[idx]
			count++
			saw = true
		} else if nums[idx] != nums[idx-1] {
			nums[count] = nums[idx]
			count++
			saw = false
		}
	}

	return count
}

// Helper fn to print these arrays
func PrintDuplicateFuncs(nums []int) {
	size := RemoveDuplicatesII(nums)

	print("[")
	for n := 0; n < size-1; n++ {
		fmt.Printf("%d, ", nums[n])
	}

	fmt.Printf("%d]\n", nums[size-1])
}

// NOTES:
// Can assume that count will be at least 1
// We can start at the 2nd element in the array since we can guarantee that the first elm is uniq
// - if nums[i] != nums[i - 1]
// 		- Can replace with last val
// - We don't need to swap cause we don't care about the rest of the array!
