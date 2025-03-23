package datatypes

import "fmt"

func TestSLinkedList() {
	root := NewLinkedList([]int{1, 2, 3, 4})
	root.Print()
	reverse := Reverse(root)

	reverse.Print()
}

// There's a good chance I'll cover a double linked list so to save potential headache,
// S will stand for Single and D will stand for Double.
// TODO: introduce generics to this struct
type SNode struct {
	Value int
	Next  *SNode
}

func NewLinkedList(nums []int) *SNode {
	if len(nums) == 0 {
		return nil
	}

	root := &SNode{
		Value: nums[0],
	}
	prevNode := root

	for i := 1; i < len(nums); i++ {
		node := SNode{
			Value: nums[i],
		}

		prevNode.Next = &node
		prevNode = &node
	}

	return root
}

func (sl *SNode) Print() {
	for sl.Next != nil {
		fmt.Printf("%d, ", sl.Value)
		sl = sl.Next
	}

	fmt.Printf("%d\n", sl.Value)
}

// Reverses the linked list of the provided
func Reverse(current *SNode) *SNode {
	// Get the next node.
	// Set current node at the back of the list
	// Repeat until next node == nil
	nextNode := current.Next
	var reverse *SNode

	for nextNode != nil {
		// Connect current node to revser
		current.Next = reverse
		// Set reverse to current node
		reverse = current
		// Assign the next node to the current node
		current = nextNode
		// Assign nextNode
		nextNode = current.Next
	}

	// We'll need to connect the last node
	current.Next = reverse
	reverse = current

	return reverse
}
