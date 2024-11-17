package flags

import (
	"flag"
	"fmt"
	"os"
)

// Flags содержит флаги командной строки.
type Flags struct {
	Command string   // Команда для выполнения
	Args    []string // Аргументы для команды
}

// ParseFlags разбирает флаги командной строки и возвращает заполненную структуру Flags.
func ParseFlags() *Flags {
	flags := &Flags{}
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "Usage: %s <command> [args...]\n", os.Args[0])
		flag.PrintDefaults()
	}
	flag.Parse()

	if flag.NArg() < 1 {
		flag.Usage()
		os.Exit(1)
	}

	flags.Command = flag.Arg(0)
	flags.Args = flag.Args()[1:]

	return flags
}
