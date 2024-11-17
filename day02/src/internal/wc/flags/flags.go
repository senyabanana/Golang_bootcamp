package flags

import (
	"flag"
)

// Flags содержит флаги командной строки.
type Flags struct {
	CountLines bool     // Флаг для подсчета строк
	CountChars bool     // Флаг для подсчета символов
	CountWords bool     // Флаг для подсчета слов
	FilePaths  []string // Пути для файлов
}

// ParseFlags разбирает флаги командной строки и возвращает заполненную структуру Flags.
func ParseFlags() *Flags {
	flags := &Flags{}
	flag.BoolVar(&flags.CountLines, "l", false, "Count lines")
	flag.BoolVar(&flags.CountChars, "m", false, "Count characters")
	flag.BoolVar(&flags.CountWords, "w", false, "Count words")
	flag.Parse()

	flags.FilePaths = flag.Args()
	return flags
}
