package main

import (
	"fmt"
	"os"

	"day02/internal/rotate/archiver"
	"day02/internal/rotate/flags"
)

func main() {
	flagSet := flags.ParseFlags()

	if err := archiver.ProcessFiles(flagSet); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
