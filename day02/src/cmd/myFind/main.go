package main

import (
	"fmt"
	"os"

	"day02/internal/find/flags"
	"day02/internal/find/walker"
)

func main() {
	checkFlags := flags.ParseFlags()
	err := walker.WalkFileSystem(checkFlags)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
}
