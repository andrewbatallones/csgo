package main

import (
	"github.com/andrewbatallones/csgo/leetcode"
)

func main() {
	leetcode.PrintDuplicateFuncs([]int{1, 1, 2, 2, 2, 3, 3, 3})
	leetcode.PrintDuplicateFuncs([]int{1, 2, 3, 4, 5})
	leetcode.PrintDuplicateFuncs([]int{1})
	leetcode.PrintDuplicateFuncs([]int{1, 2, 3, 3, 3, 3, 3, 3, 4, 5})
}
