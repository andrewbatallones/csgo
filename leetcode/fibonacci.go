package leetcode

import "fmt"

func fib(n int) int {
	if n <= 1 {
		return n
	}

	// Init the memory
	mem := make(map[int]int)

	// We will assign the left/right values to memory.
	// As we recurse into each n-k value, we'll save it to memory and can recall it.
	mem[n-1] = fibMem(n-1, &mem)
	mem[n-2] = fibMem(n-2, &mem)

	// We don't need to recruse here since we have already calculated the values.
	return mem[n-1] + mem[n-2]
}

func fibMem(n int, mem *map[int]int) int {
	if n <= 1 {
		return n
	}

	// Pull and check if the left value exists
	left, lExist := (*mem)[n-1]
	if !lExist {
		// We'll need to calculate them if it doesn't.
		// Pass in the memory to make sure it saves accross each recursive func.
		(*mem)[n-1] = fibMem(n-1, mem)
		left = (*mem)[n-1]
	}

	// Same here.
	right, rExist := (*mem)[n-2]
	if !rExist {
		(*mem)[n-2] = fibMem(n-2, mem)
		right = (*mem)[n-2]
	}

	// Return the 2 added values.
	return left + right
}

func TestFib() {
	fmt.Printf("%d\n", fib(2))
	fmt.Printf("%d\n", fib(3))
	fmt.Printf("%d\n", fib(4))
	fmt.Printf("%d\n", fib(10))
}
