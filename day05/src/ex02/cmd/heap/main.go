package main

import (
	"fmt"
	"log"

	present "binary-tree/pkg/present-heap"
)

func main() {
	presents := []present.Present{
		{Value: 5, Size: 1},
		{Value: 4, Size: 5},
		{Value: 3, Size: 1},
		{Value: 5, Size: 2},
	}

	presentHeap := present.NewPresentHeap(presents)

	n := 2
	coolestPresents, err := presentHeap.GetNCoolestPresents(n)
	if err != nil {
		log.Println("error:", err)
	} else {
		fmt.Println("Coolest Presents:", coolestPresents)
	}
}
