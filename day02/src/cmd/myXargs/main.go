package main

import (
	"fmt"
	"os"

	"day02/internal/xargs/flags"
	"day02/internal/xargs/runner"
)

func main() {
	flagSet := flags.ParseFlags()

	if flagSet.Command == "" {
		fmt.Println("Usage: ./myXargs <command> [args...]")
		os.Exit(1)
	}

	if err := runner.ProcessStdin(flagSet.Command, flagSet.Args); err != nil {
		fmt.Fprintf(os.Stderr, "Error processing stdin: %v\n", err)
		os.Exit(1)
	}
}
