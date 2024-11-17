package runner

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"sync"
)

// RunCommand выполняет команду с переданными аргументами.
func runCommand(command string, args []string) error {
	cmd := exec.Command(command, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

// ProcessStdin читает стандартный ввод и выполняет команду для каждой строки ввода.
func ProcessStdin(command string, args []string) error {
	scanner := bufio.NewScanner(os.Stdin)
	var wg sync.WaitGroup

	for scanner.Scan() {
		line := scanner.Text()
		wg.Add(1)
		go func(line string) {
			defer wg.Done()
			// объединяем аргументы
			fullArgs := append(args, strings.Fields(line)...)
			if err := runCommand(command, fullArgs); err != nil {
				fmt.Fprintf(os.Stderr, "Error running command: %v\n", err)
			}
		}(line)
	}

	wg.Wait()
	return scanner.Err()
}
