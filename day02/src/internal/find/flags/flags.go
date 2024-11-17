package flags

import (
	"flag"
	"fmt"
	"os"
)

// Flags содержит флаги командной строки.
type Flags struct {
	ShowFiles    bool   // Показывать файлы
	ShowDirs     bool   // Показывать директории
	ShowSymlinks bool   // Показывать символические ссылки
	Ext          string // Фильтр по расширению файлов (работает только с -f)
	Path         string // Путь для поиска
}

// ParseFlags разбирает флаги командной строки и возвращает заполненную структуру Flags.
func ParseFlags() *Flags {
	flags := &Flags{}
	flag.BoolVar(&flags.ShowFiles, "f", false, "Show files")
	flag.BoolVar(&flags.ShowDirs, "d", false, "Show directories")
	flag.BoolVar(&flags.ShowSymlinks, "sl", false, "Show symlinks")
	flag.StringVar(&flags.Ext, "ext", "", "Filter files by extension (only works with -f)")
	flag.Parse()

	if flag.NArg() < 1 {
		fmt.Println("Usage: ./myFind [options] <path>")
		os.Exit(1)
	}

	flags.Path = flag.Arg(0)
	return flags
}
