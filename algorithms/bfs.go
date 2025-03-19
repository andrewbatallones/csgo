package algorithms

import (
	"fmt"
	"slices"
)

// Breath First Search
// Essentially, visit all 1st children before searching next level
// Ideal for shortest path problems

type Coord struct {
	Row int
	Col int
}

// This is a Flood Fill problem
func PaintFill(img [][]int, row, col, p int) {
	var coord Coord
	start := img[row][col]
	queue := []*Coord{{Row: row, Col: col}}
	visited := []*Coord{}

	for len(queue) > 0 {
		coord, queue = pop(queue)
		visited = append(visited, &coord)

		// Assign new val
		img[coord.Row][coord.Col] = p

		neighbors := coord.getNeighbors(img, start)

		for _, nCoord := range neighbors {
			if !slices.Contains(visited, &nCoord) {
				queue = append(queue, &nCoord)
			}
		}
	}
}

func (c *Coord) getNeighbors(img [][]int, start int) []Coord {
	left := Coord{
		Row: c.Row - 1,
		Col: c.Col,
	}
	right := Coord{
		Row: c.Row + 1,
		Col: c.Col,
	}
	up := Coord{
		Row: c.Row,
		Col: c.Col - 1,
	}
	down := Coord{
		Row: c.Row,
		Col: c.Col + 1,
	}

	indices := []Coord{left, right, up, down}
	ret := []Coord{}

	for _, coord := range indices {
		if coord.isValid(img, start) {
			ret = append(ret, coord)
		}
	}

	return ret
}

func (c *Coord) isValid(img [][]int, start int) bool {
	return c.Row >= 0 && c.Col >= 0 && c.Row < len(img) && c.Col < len(img) && img[c.Row][c.Col] == start
}

func pop(sl []*Coord) (Coord, []*Coord) {
	result := sl[0]
	sl = sl[1:]

	return *result, sl
}

func TestPaintFill() {
	img := [][]int{
		{1, 0, 1, 1, 0},
		{0, 1, 0, 1, 0},
		{1, 1, 1, 1, 1},
		{0, 0, 1, 0, 1},
		{1, 0, 0, 0, 0},
	}

	printImg(img)

	PaintFill(img, 2, 2, 2)

	printImg(img)
}

func printImg(img [][]int) {
	for _, row := range img {
		fmt.Print("[ ")

		for _, cell := range row {
			fmt.Printf("%d ", cell)
		}

		fmt.Print("]\n")
	}
}
