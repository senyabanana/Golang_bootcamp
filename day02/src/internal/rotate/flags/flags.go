package flags

import (
	"flag"
	"fmt"
	"os"
)

// Flags содержит флаги командной строки.
type Flags struct {
	ArchiveDir string   // Директория для архивации файлов
	FilePaths  []string // Пути к файлам логов
}

// ParseFlags разбирает флаги командной строки и возвращает заполненную структуру Flags.
func ParseFlags() *Flags {
	flags := &Flags{}
	flag.StringVar(&flags.ArchiveDir, "a", "", "Directory to store archived logs")
	flag.Parse()

	if flag.NArg() < 1 {
		fmt.Println("Usage: ./myRotate [options] <file1> <file2> ...")
		os.Exit(1)
	}

	flags.FilePaths = flag.Args()
	return flags
}
