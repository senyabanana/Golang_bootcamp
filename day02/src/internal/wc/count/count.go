package count

import (
	"bufio"
	"os"
	"sync"

	"day02/internal/wc/flags"
)

// Result представляет результат подсчета для одного файла.
type Result struct {
	FilePath string
	Count    int
}

// countLines - функция для подсчета строк.
func countLines(file *os.File) (int, error) {
	scanner := bufio.NewScanner(file)
	lines := 0
	for scanner.Scan() {
		lines++
	}

	return lines, scanner.Err()
}

// countChars - функция для подсчета символов.
func countChars(file *os.File) (int, error) {
	scanner := bufio.NewScanner(file)
	chars := 0
	for scanner.Scan() {
		chars += len(scanner.Text())
	}

	return chars, scanner.Err()
}

// countWords - функция для подсчета слов.
func countWords(file *os.File) (int, error) {
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	words := 0
	for scanner.Scan() {
		words++
	}

	return words, scanner.Err()
}

// ProcessFile обрабатывает один файл в зависимости от заданных флагов.
func processFile(filepath string, flags *flags.Flags) (*Result, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var count int
	if flags.CountLines {
		count, err = countLines(file)
	} else if flags.CountChars {
		count, err = countChars(file)
	} else if flags.CountWords {
		count, err = countWords(file)
	} else {
		count, err = countWords(file)
	}

	if err != nil {
		return nil, err
	}

	return &Result{
		FilePath: filepath,
		Count:    count,
	}, nil
}

// ProcessFiles обрабатывает несколько файлов параллельно.
func ProcessFiles(filePaths []string, flags *flags.Flags) ([]*Result, error) {
	var wg sync.WaitGroup
	results := make([]*Result, len(filePaths))
	errs := make([]error, len(filePaths))

	for i, filePath := range filePaths {
		wg.Add(1)
		go func(i int, filePath string) {
			defer wg.Done()
			result, err := processFile(filePath, flags)
			results[i] = result
			errs[i] = err
		}(i, filePath)
	}

	wg.Wait()

	for _, err := range errs {
		if err != nil {
			return nil, err
		}
	}

	return results, nil
}
