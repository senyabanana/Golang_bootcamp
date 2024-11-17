package data

import (
	"encoding/csv"
	"fmt"
	"os"
)

// ReadCSVData читает данные из CSV файла и возвращает их в виде двумерного среза строк.
func ReadCSVData(filePath string) ([][]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.Comma = '\t' // Определяет разделитель колонок
	lines, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	if len(lines) < 1 {
		return nil, fmt.Errorf("No data found in %s", filePath)
	}

	return lines, nil
}
