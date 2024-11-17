package data

import (
	"encoding/json"
	"os"
)

// PropertiesType представляет тип свойства в маппинге.
type PropertiesType struct {
	Type string `json:"type"`
}

// Properties представляет свойства схемы индекса.
type Properties struct {
	Id       PropertiesType `json:"id"`
	Name     PropertiesType `json:"name"`
	Address  PropertiesType `json:"address"`
	Phone    PropertiesType `json:"phone"`
	Location PropertiesType `json:"location"`
}

// Schema представляет схему индекса.
type Schema struct {
	Properties Properties `json:"properties"`
}

// IndexSettings представляет настройки индекса, включая схему.
type IndexSettings struct {
	Mappings Schema `json:"mappings"`
}

// ReadMapping читает схему индекса из JSON файла.
func ReadMapping(filePath string) (Schema, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return Schema{}, err
	}
	defer file.Close()

	var mapping Schema
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&mapping); err != nil {
		return Schema{}, err
	}

	return mapping, nil
}
