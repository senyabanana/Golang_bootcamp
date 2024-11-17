package main

import (
	"fmt"

	present "binary-tree/pkg/present-heap"
)

func main() {
	presents := []present.Present{
		{Value: 5, Size: 3},
		{Value: 7, Size: 4},
		{Value: 2, Size: 2},
		{Value: 9, Size: 5},
	}
	capacity := 9

	pc := present.NewPresentCollection(presents)
	selectedPresents := pc.GrabPresents(capacity)
	fmt.Println("Selected Presents:")
	for _, p := range selectedPresents {
		fmt.Printf("Value: %d, Size: %d\n", p.Value, p.Size)
	}
}
