package algorithms

// Merge Sort
// The sort splits the array until it cannot be split anymore,
// then it goes back up merging them in the correct order
func MergeSort(arr []int) []int {
	// We can stop at 2 length since comparing them should return a sorted slice
	if len(arr) < 2 {
		return arr
	}

	mid := len(arr) / 2

	// Here, we are recursing until we hit a the smallest array
	left := MergeSort(arr[:mid])
	right := MergeSort(arr[mid:])

	return merge(left, right, len(arr))
}

// Merges the 2 slices sorting them.
// It also uses the length (since the func knows how long the combined slices are)
// Another option is to just pass the slices and construct the new slice by adding
// both lengths
func merge(left, right []int, length int) []int {
	arr := make([]int, length)
	i, j, k := 0, 0, 0

	for i < len(left) && j < len(right) {
		if left[i] > right[j] {
			arr[k] = right[j]
			j++
		} else {
			arr[k] = left[i]
			i++
		}

		k++
	}

	for i < len(left) {
		arr[k] = left[i]
		k++
		i++
	}

	for j < len(right) {
		arr[k] = right[j]
		k++
		j++
	}

	return arr
}
