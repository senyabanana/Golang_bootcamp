package dbreader

import (
	"fmt"
	"io"
	"os"
	"strings"
)

// DBReader - интерфейс для чтения и конвертации баз данных рецептов.
type DBReader interface {
	Read(r io.Reader) (*Recipes, error)
	Convert(recipes *Recipes) ([]byte, error)
}

// Ingredient представляет ингредиент рецепта.
type Ingredient struct {
	Name    string `json:"ingredient_name" xml:"itemname"`
	Count   string `json:"ingredient_count" xml:"itemcount"`
	Unit    string `json:"ingredient_unit,omitempty" xml:"itemunit,omitempty"`
	Comment string `json:"-" xml:",comment"`
}

// Cake представляет торт с его параметрами.
type Cake struct {
	Name        string       `json:"name" xml:"name"`
	Time        string       `json:"time" xml:"stovetime"`
	Ingredients []Ingredient `json:"ingredients" xml:"ingredients>item"`
	Comment     string       `json:"-" xml:",comment"`
}

// Recipes представляет коллекцию тортов.
type Recipes struct {
	Cakes   []Cake `json:"cake" xml:"cake"`
	Comment string `json:"-" xml:",comment"`
}

// ReadDB читает базу данных из файла по указанному пути и возвращает структуру Recipes.
// Определяет формат файла (JSON или XML) и использует соответствующий ридер для чтения данных.
// Возвращает указатель на Recipes, DBReader для дальнейшей обработки и ошибку, если таковая возникла.
func ReadDB(filepath string) (*Recipes, DBReader, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, nil, fmt.Errorf("не удалось открыть файл: %v", err)
	}
	defer file.Close()

	var reader DBReader
	if strings.HasSuffix(filepath, ".json") {
		reader = JSONReader{}
	} else if strings.HasSuffix(filepath, ".xml") {
		reader = XMLReader{}
	} else {
		return nil, nil, fmt.Errorf("неподдерживаемый формат файла. Пожалуйста, предоставьте файл в формате .json или .xml")
	}

	recipes, err := reader.Read(file)
	if err != nil {
		return nil, nil, fmt.Errorf("не удалось прочитать файл: %v", err)
	}

	return recipes, reader, nil
}
