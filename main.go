package main

import (
	"github.com/andrewbatallones/csgo/algorithms"
	"github.com/andrewbatallones/csgo/leetcode"
)

func main() {
	sorted := algorithms.MergeSort([]int{2, 4, 6, 8, 1, 3, 5})

	leetcode.PrintSlice(sorted)
}
