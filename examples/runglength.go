package examples

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// This will compress a string.
// Example: aaaabbccca => 4a2b3c1a
func Runlength(input string) string {
	compressed := ""

	if len(input) > 0 {
		// Want to grab the string and the counter.
		count := 1
		letterByte := input[0]

		for i := 1; i < len(input); i++ {
			// Check if letter is the same
			// If not -> add count/letter to compressed
			// 		- reset count/letter
			if input[i] != letterByte {
				compressed += fmt.Sprintf("%d%s", count, string(letterByte))
				count = 1
				letterByte = input[i]
			} else {
				count++
			}
		}

		// Need to add the last letter
		compressed += fmt.Sprintf("%d%s", count, string(letterByte))
	}

	return compressed
}

// be going backwards: 4a2b3c1a => aaaabbccca
func DecodeRun(input string) (string, error) {
	// Know that it will be in pairs of 2
	// First check if input is blank
	if len(input) == 0 {
		return "", nil
	// Check if it is not divisable by 2
	} else if len(input)%2 != 0 {
		err := errors.New("unable to decode compression string, corrupted input.")
		return "", err
	} else {
		output := ""

		// increment by 2 (since it's a number, letter pair)
		for i := 0; i < len(input); i += 2 {
			// cast from string to int
			times, err := strconv.Atoi(string(input[i]))
			if err != nil {
				return "", errors.New("cannot decode string")
			}
			letter := string(input[i+1])

			output += strings.Repeat(letter, times)
		}

		return output, nil
	}
}

// NOTES:
// - learned that getting the index of a string will return the bytes (need to cast it back to string)
// - It looks like for an interview, testing means that they want you to walk through the code manually to see if you absolutely understand what you wrote
