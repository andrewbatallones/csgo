package leetcode

import (
	"strings"
	"unicode"
)

// 5. Longest Palindromic Substring
// Return the longest palindrome substring within the string
func LongestPalindrome(s string) string {
	if len(s) <= 1 {
		return s
	}

	maxLen := 1
	maxStr := string(s[0])

	// Make a 2D memory of the length of the word
	mem := make([][]bool, len(s))
	for i := range mem {
		mem[i] = make([]bool, len(s))
	}

	for idx := range s {
		mem[idx][idx] = true

		for j := range idx {
			if s[j] == s[idx] && (idx-j <= 2 || mem[j+1][idx-1]) {
				mem[j][idx] = true

				if idx-j+1 > maxLen {
					maxLen = idx - j + 1
					maxStr = s[j : idx+1]
				}
			}
		}
	}

	return maxStr
}

// 125. Valid Palindrome
// Check if lowercase string is a palindrome
func IsPalindrome(s string) bool {
	// downcase the string first
	// have 2 pointers (beginning, end)
	// iterate through each pointer
	// 		skip to next pointer if it is not a alphanumeric character
	// 		return `false` if characters are not equal
	s = strings.ToLower(s)
	left := 0
	right := len(s) - 1

	for left <= right {
		if !isAlphanumeric(rune(s[left])) {
			left++
		} else if !isAlphanumeric(rune(s[right])) {
			right--
		} else if s[left] != s[right] {
			return false
		} else {
			left++
			right--
		}
	}

	return true
}

func isAlphanumeric(r rune) bool {
	return unicode.IsLetter(r) || unicode.IsDigit(r)
}
