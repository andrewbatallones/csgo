package leetcode

func MajorityElement(nums []int) int {
	numCount := make(map[int]int)
	result := 0

	for i := 0; i < len(nums); i++ {
		numCount[nums[i]]++
	}

	for idx, val := range numCount {
		if numCount[result] < val {
			result = idx
		}
	}

	return result
}
