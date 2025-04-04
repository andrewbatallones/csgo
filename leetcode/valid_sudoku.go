package leetcode

import (
	"fmt"
	"slices"
)

// 36. Valid Sudoku
func ValidSudoku(board [][]byte) bool {
	// Iterate through each row, then each column (that should get each cell)
	// Want to create a map that keeps track of each row and column
	cols := make(map[int][]byte)
	rows := make(map[int][]byte)
	squares := make(map[string][]byte)

	for rowIdx, row := range board {
		for colIdx, cell := range row {
			// This just checks if the cell is equal to []byte(".")
			if cell == 46 {
				continue
			}

			grid := fmt.Sprintf("%d-%d", rowIdx/3, colIdx/3)

			// Check row if number already exists
			// Check col if number already exists
			// Check grid if number already exists
			if slices.Contains(rows[rowIdx], cell) ||
				slices.Contains(cols[colIdx], cell) ||
				slices.Contains(squares[grid], cell) {
				return false
			}

			// Add them to seen
			rows[rowIdx] = append(rows[rowIdx], cell)
			cols[colIdx] = append(cols[colIdx], cell)
			squares[grid] = append(squares[grid], cell)
		}
	}

	return true
}
