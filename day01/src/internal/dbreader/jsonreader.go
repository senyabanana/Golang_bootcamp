package dbreader

import (
	"encoding/json"
	"encoding/xml"
	"io"
)

// JSONReader реализует интерфейс DBReader для чтения JSON файлов.
type JSONReader struct{}

// Read читает данные из io.Reader и демаршалит их в структуру Recipes.
// Возвращает указатель на Recipes и ошибку, если таковая возникла.
func (jr JSONReader) Read(r io.Reader) (*Recipes, error) {
	var recipes Recipes
	decoder := json.NewDecoder(r)
	if err := decoder.Decode(&recipes); err != nil {
		return nil, err
	}

	return &recipes, nil
}

// Convert принимает структуру Recipes и маршалит её в XML формат.
// Возвращает XML данные в виде байтов и ошибку, если таковая возникла.
func (jr JSONReader) Convert(recipes *Recipes) ([]byte, error) {
	xmlData, err := xml.MarshalIndent(recipes, "", "\t")
	if err != nil {
		return nil, err
	}

	return xmlData, nil
}
