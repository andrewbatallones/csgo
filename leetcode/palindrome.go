package leetcode

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
