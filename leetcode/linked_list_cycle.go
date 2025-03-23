package leetcode

import (
	"slices"

	"github.com/andrewbatallones/csgo/datatypes"
)

// 141. Linked List Cycle
// Checks to see if the given head has a cycle (doesn't have a nil tail)
func HasCycle(head *datatypes.SNode) bool {
	visited := []*datatypes.SNode{}

	// This is if we were given an empty linked list
	if head == nil {
		return false
	}

	for head.Next != nil {
		if slices.Contains(visited, head) {
			return true
		}

		// Add head to visted
		visited = append(visited, head)
		// Set head to next node
		head = head.Next
	}

	return false
}

// Another solution that people have found is a slow/fast pointer.
// 1 pointer moves 1 spaces, another pointer moves 2 spaces.
// Eventually, they will meet up which will signify a cycle.
