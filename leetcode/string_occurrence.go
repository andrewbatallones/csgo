package leetcode

// 28. Find the index of the First Occurrence in a String
func StrStr(haystack, needle string) int {
	// We are just comparing all strings within the haystack
	for idx := 0; idx+len(needle)-1 < len(haystack); idx++ {
		if needle == haystack[idx:idx+len(needle)] {
			return idx
		}
	}

	return -1
}
