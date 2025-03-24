package leetcode

import "fmt"

// 200. Number of Islands
// Given a m x n grid of '1' and '0', return the number of islands ('1')
func NumIslands(grid [][]byte) int {
	// Essentially, we're going to iterate through the entire grid.
	// If we have a 1, we will perform a DFS (Depth First Search) to iterate
	// through the island.
	// As we traverse, we append/remove from the stack.
	// As we backtrack, we just set the coord to '0' (this will prevent double counting).
	// when done, increment count++
	count := 0

	// This will iterate through each cell
	for m := 0; m < len(grid); m++ {
		for n := 0; n < len(grid[m]); n++ {
			if grid[m][n] == '1' {
				dfs(grid, m, n)
				count++
			}
		}
	}

	return count
}

// Performs the dfs and sets byte to '0'
func dfs(grid [][]byte, m, n int) {
	// Check if out of bounds
	// Also need to check if we hit a water
	if m < 0 || n < 0 || m >= len(grid) || n >= len(grid[m]) || grid[m][n] == '0' {
		return
	}

	grid[m][n] = '0'

	dfs(grid, m+1, n) // right
	dfs(grid, m, n+1) // down
	dfs(grid, m-1, n) // left
	dfs(grid, m, n-1) // up
}

func TestIslandCount() {
	grid := [][]byte{
		{'1', '1', '1', '1', '0'},
		{'1', '1', '0', '1', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '0', '0', '0'},
	}

	fmt.Printf("%d\n", NumIslands(grid))
}
