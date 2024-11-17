package main

import (
	"flag"
	"fmt"
	"json-xml-basics/internal/fscomparer"
)

func main() {
	oldFilePath := flag.String("old", "", "Path to the old filesystem snapshot")
	newFilePath := flag.String("new", "", "Path to the new filesystem snapshot")
	flag.Parse()

	if *oldFilePath == "" || *newFilePath == "" {
		fmt.Println("Пожалуйста, укажите пути как к старому, так и к новому файлам файловой системы")
		return
	}

	oldData, err := fscomparer.ReadFile(*oldFilePath)
	if err != nil {
		fmt.Printf("Ошибка при обработке старого файла: %v\n", err)
		return
	}

	newData, err := fscomparer.ReadFile(*newFilePath)
	if err != nil {
		fmt.Printf("Ошибка при обработке нового файла: %v\n", err)
		return
	}

	fscomparer.CompareFS(oldData, newData)
}
