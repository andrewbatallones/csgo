package main

import (
	"github.com/andrewbatallones/csgo/leetcode"
)

func main() {
	tn1 := leetcode.NewTree([]int{1, 2, 3})
	tn2 := leetcode.NewTree([]int{1, 2, 3})

	leetcode.IsSameTree(tn1, tn2)
}
