package leetcode

import (
	"fmt"
)

// Converts Roman to an int
// I: 1,
// V: 5,
// X: 10,
// L: 50,
// C: 100,
// D: 500,
// M: 1000,
// I can be before V/X => 4/9
// X can be before L/C => 40/90
// C can be before D/M => 400/900
func RomantoInt(s string) int {
	result := 0
	// Iterate through the string.
	// Want to keep track of what letter is before.
	// Only need to check if the prev is lower.
	for idx, hex := range s {
		// Get letter and prev letter (only if it can be a previous, or just set it to the default "")
		letter := string(hex)
		prevLetter := ""

		if idx > 0 {
			prevLetter = string(s[idx-1])
		}

		// Get values
		letterVal := returnInt(letter)
		prevLetterVal := returnInt(prevLetter)

		// Check if prev is lower than current
		// We don't care about the 0 val since subtracting 0 does nothing
		if letterVal > prevLetterVal {
			// Want to subtract prevLetterVal twice because we need to null the first time
			// it was added.
			result += letterVal - prevLetterVal - prevLetterVal
		} else {
			result += letterVal
		}
	}

	return result
}

func returnInt(s string) int {
	switch s {
	case "I":
		return 1
	case "V":
		return 5
	case "X":
		return 10
	case "L":
		return 50
	case "C":
		return 100
	case "D":
		return 500
	case "M":
		return 1000
	default:
		return 0
	}
}

func returnRoman(i int) (string, int) {
	if i >= 1000 {
		return "M", 1000
	} else if i >= 900 {
		return "CM", 900
	} else if i >= 500 {
		return "D", 500
	} else if i >= 400 {
		return "CD", 400
	} else if i >= 100 {
		return "C", 100
	} else if i >= 90 {
		return "XC", 90
	} else if i >= 50 {
		return "L", 50
	} else if i >= 40 {
		return "XL", 40
	} else if i >= 10 {
		return "X", 10
	} else if i >= 9 {
		return "IX", 9
	} else if i >= 5 {
		return "V", 5
	} else if i >= 4 {
		return "IV", 4
	} else if i >= 1 {
		return "I", 1
	} else {
		return "", 0
	}
}

func TestRomantoInt() {
	fmt.Printf("%d\n", RomantoInt("III"))
	fmt.Printf("%d\n", RomantoInt("LVIII"))
	fmt.Printf("%d\n", RomantoInt("MCMXCIV"))
	fmt.Printf("%d\n", RomantoInt("XL"))
}

func InttoRoman(num int) string {
	result := ""

	// Keep removing until num is empty
	for num > 0 {
		roman, val:= returnRoman(num)

		result += roman
		num -= val
	}

	return result
}

func TestInttoRoman() {
	fmt.Printf("%s\n", InttoRoman(3749))
	fmt.Printf("%s\n", InttoRoman(58))
	fmt.Printf("%s\n", InttoRoman(1994))
	fmt.Printf("%s\n", InttoRoman(40))
}
