package main

import (
	"ex00/internal/data"
	elastic "ex00/internal/elasticsearch"
	"github.com/sirupsen/logrus"
)

const (
	indexName = "places"
	jsonData  = `data/schema.json`
	csvData   = `data/data.csv`
)

func main() {
	// Создаем клиента Elasticsearch
	es, err := elastic.NewClient()
	if err != nil {
		logrus.Fatalf("Error creating Elasticsearch client: %s\n", err)
	}

	// Читаем схему индекса
	mapping, err := data.ReadMapping(jsonData)
	if err != nil {
		logrus.Fatalf("Error reading file: %s\n", err)
	}

	// Создаем и настраиваем индекс
	elastic.CreateAndMapping(es, indexName, mapping)

	// Читаем данные из CSV
	lines, err := data.ReadCSVData(csvData)
	if err != nil {
		logrus.Fatalf("Error reading data.csv: %s", err)
	}

	// Создаем bulk-запрос
	bulkRequest := elastic.CreateBulkRequest(lines)

	// Отправляем bulk-запрос
	elastic.SendBulkRequest(es, indexName, bulkRequest)
}
