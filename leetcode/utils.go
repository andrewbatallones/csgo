package leetcode

import "fmt"

func MoveBack(nums []int, idx int) {
	for idx < len(nums)-1 {
		Swap(nums, idx, idx+1)
		idx++
	}
}

func Swap(nums []int, left, right int) {
	tmp := nums[left]
	nums[left] = nums[right]
	nums[right] = tmp
}

func PrintSlice(nums []int) {
	print("[")
	for n := 0; n < len(nums)-1; n++ {
		fmt.Printf("%d, ", nums[n])
	}

	fmt.Printf("%d]\n", nums[len(nums)-1])
}
