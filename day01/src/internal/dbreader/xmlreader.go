package dbreader

import (
	"encoding/json"
	"encoding/xml"
	"io"
)

// XMLReader реализует интерфейс DBReader для чтения XML файлов.
type XMLReader struct{}

// Read читает данные из io.Reader и демаршалит их в структуру Recipes.
// Возвращает указатель на Recipes и ошибку, если таковая возникла.
func (xr XMLReader) Read(r io.Reader) (*Recipes, error) {
	var recipes Recipes
	decoder := xml.NewDecoder(r)
	if err := decoder.Decode(&recipes); err != nil {
		return nil, err
	}

	return &recipes, nil
}

// Convert принимает структуру Recipes и маршалит её в JSON формат.
// Возвращает JSON данные в виде байтов и ошибку, если таковая возникла.
func (xr XMLReader) Convert(recipes *Recipes) ([]byte, error) {
	jsonData, err := json.MarshalIndent(recipes, "", "\t")
	if err != nil {
		return nil, err
	}

	return jsonData, nil
}
