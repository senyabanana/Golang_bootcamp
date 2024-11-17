package main

import (
	"flag"
	"fmt"
	
	"json-xml-basics/internal/dbcomparer"
	"json-xml-basics/internal/dbreader"
)

func main() {
	oldFilePath := flag.String("old", "", "Path to the original database file")
	newFilePath := flag.String("new", "", "Path to the new database file")
	flag.Parse()

	if *oldFilePath == "" || *newFilePath == "" {
		fmt.Println("Пожалуйста, укажите пути как к старым, так и к новым файлам базы данных")
		return
	}

	oldRecipes, _, err := dbreader.ReadDB(*oldFilePath)
	if err != nil {
		fmt.Printf("Не удалось обработать файл: %v\n", err)
		return
	}

	newRecipes, _, err := dbreader.ReadDB(*newFilePath)
	if err != nil {
		fmt.Printf("Не удалось обработать файл: %v\n", err)
		return
	}

	dbcomparer.CompareDB(oldRecipes, newRecipes)
}
