package main

import "github.com/andrewbatallones/csgo/datatypes"

func main() {
	mHeap := datatypes.NewMinHeap([]int{5, 4, 3, 2, 1})

	mHeap.Print()
}
