package fscomparer

import (
	"bufio"
	"fmt"
	"os"
)

// ReadFile читает файл построчно и сохраняет каждую строку в хэш-таблицу.
func ReadFile(filepath string) (map[string]struct{}, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, fmt.Errorf("не удалось открыть файл: %v", err)
	}
	defer file.Close()

	fileSet := make(map[string]struct{})
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fileSet[scanner.Text()] = struct{}{}
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("Ошибка чтения файла: %v\n", err)
	}

	return fileSet, nil
}

// CompareFS сравнивает два набора файлов и выводит добавленные и удаленные файлы.
func CompareFS(oldSet, newSet map[string]struct{}) {
	for newFile := range newSet {
		if _, exists := oldSet[newFile]; !exists {
			fmt.Printf("ADDED %s\n", newFile)
		}
	}
	for oldFile := range oldSet {
		if _, exists := newSet[oldFile]; !exists {
			fmt.Printf("REMOVED %s\n", oldFile)
		}
	}
}
