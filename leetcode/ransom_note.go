package leetcode

import "fmt"

// 383. Ransom Note
// Given a note, and letters
func RandomNote(ransomNote, magazine string) bool {
	// Didn't realize rune is a type
	letters := make(map[rune]int)

	// Add letters to map
	for _, rune := range magazine {
		letters[rune]++
	}

	// Check letters within note if it's able to construct note
	// Delete each letter from map
	for _, ransomRune := range ransomNote {
		if letters[ransomRune] <= 0 {
			return false
		}

		letters[ransomRune]--
	}

	return true
}

func TestRansomNote() {
	fmt.Printf("%v\n", RandomNote("a", "b"))
	fmt.Printf("%v\n", RandomNote("aa", "ab"))
	fmt.Printf("%v\n", RandomNote("aa", "aab"))
	fmt.Printf("%v\n", RandomNote("aa", "aab"))
}
