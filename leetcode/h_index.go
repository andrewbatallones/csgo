package leetcode

import (
	"fmt"
	"sort"
)

// 274. H-Index
// calculates the h-index of given citations
// `paper` is the paper published
// `citations[paper]` is how many citations that paper got
func HIndex(citations []int) int {
	hIndex := 0

	// Reverse sort (most -> least)
	// Note: The func params are indicies
	sort.Slice(citations, func(i, j int) bool {
		return citations[j] < citations[i]
	})

	// Go through each citation w/ index as the paper
	for paper, citation := range citations {
		// Want to check paper (n'th paper) and if it's smaller than the h-index, we raise the h-index.
		if citation < paper+1 {
			break
		}

		hIndex++
	}

	return hIndex
}

func TestHIndex() {
	fmt.Printf("%d\n", HIndex([]int{3, 0, 6, 1, 5}))
	fmt.Printf("%d\n", HIndex([]int{1, 3, 1}))
	fmt.Printf("%d\n", HIndex([]int{5}))
	fmt.Printf("%d\n", HIndex([]int{5, 1}))
	fmt.Printf("%d\n", HIndex([]int{1, 5}))
	fmt.Printf("%d\n", HIndex([]int{0}))
	fmt.Printf("%d\n", HIndex([]int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1, 10}))
	fmt.Printf("%d\n", HIndex([]int{0, 0, 0, 0, 0, 0, 0, 0, 100}))
}
