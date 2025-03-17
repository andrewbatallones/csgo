package leetcode


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
