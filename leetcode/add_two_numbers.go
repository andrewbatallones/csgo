package leetcode

import (
	"github.com/andrewbatallones/csgo/datatypes"
)

// 2. Add Two Numbers
// Each list will guarantee to have at least 1 node
// values are from 0 to 9
func AddTwoNumbers(l1, l2 *datatypes.SNode) *datatypes.SNode {
	num, carryOver := addTens(l1.Value, l2.Value)
	head := &datatypes.SNode{
		Value: num,
	}
	current := head

	l1 = l1.Next
	l2 = l2.Next

	// Add until both lists are empty
	for l1 != nil || l2 != nil {
		var (
			num1 int
			num2 int
		)

		if l1 != nil {
			num1 = l1.Value
			l1 = l1.Next
		}

		if l2 != nil {
			num2 = l2.Value
			l2 = l2.Next
		}

		// We need to add in the carryOver number as well (in case it comes out to be a 10 again)
		num, carryOver = addTens(num1+carryOver, num2)
		node := &datatypes.SNode{
			Value: num,
		}

		current.Next = node
		current = node
	}

	// If there's a carryover, we'll need to create a new node for it
	if carryOver > 0 {
		node := &datatypes.SNode{
			Value: carryOver,
		}

		current.Next = node
	}

	return head
}

// Adds the 2 numbers. Will also include the carryover if needed
func addTens(num1, num2 int) (int, int) {
	result := num1 + num2

	if result >= 10 {
		return result % 10, 1
	}

	return result, 0
}

func TestAddTwoNumbers() {
	l1 := datatypes.NewLinkedList([]int{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1})
	l2 := datatypes.NewLinkedList([]int{5, 6, 4})
	sum := AddTwoNumbers(l1, l2)
	sum.Print()
}
