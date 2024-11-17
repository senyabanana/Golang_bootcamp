package main

import (
	"flag"
	"fmt"

	"json-xml-basics/internal/dbreader"
)

func main() {
	filePath := flag.String("f", "", "path to the database file")
	flag.Parse()

	if *filePath == "" {
		fmt.Println("Пожалуйста, укажите путь к файлу базы данных, используя флаг -f.")
		return
	}

	recipes, reader, err := dbreader.ReadDB(*filePath)
	if err != nil {
		fmt.Printf("Не удалось обработать файл: %v\n", err)
		return
	}

	output, err := reader.Convert(recipes)
	if err != nil {
		fmt.Printf("Не удалось замаршаллить: %v\n", err)
		return
	}
	fmt.Println(string(output))
}
