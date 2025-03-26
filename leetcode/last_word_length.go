package leetcode

import "strings"

// 58. Length of Last Word
// find the length of the last **non spaced** word
func LengthOfLastWord(s string) int {
	// Truncate word (basically remove all trailing spaces)
	truncateIdx := len(s) - 1

	for truncateIdx >= 0 {
		if string(s[truncateIdx]) != " " {
			break
		}

		truncateIdx--
	}

	truncateS := s[:truncateIdx+1]

	// Split word by space
	words := strings.Split(truncateS, " ")

	// Get len(words[len(words)-1])
	return len(words[len(words)-1])
}
