package main

import (
	"fmt"
	"os"
	"path/filepath"

	"day02/internal/wc/count"
	"day02/internal/wc/flags"
)

func main() {
	flagSet := flags.ParseFlags()

	if len(flagSet.FilePaths) == 0 {
		fmt.Println("Usage: ./myWc [options] <file1> <file2> ...")
		os.Exit(1)
	}

	results, err := count.ProcessFiles(flagSet.FilePaths, flagSet)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error processing files: %v\n", err)
		os.Exit(1)
	}

	// Вывод результатов
	for _, result := range results {
		if result != nil {
			fmt.Printf("%d\t%s\n", result.Count, filepath.Base(result.FilePath))
		}
	}
}
