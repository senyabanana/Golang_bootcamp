package elasticsearch

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"ex00/internal/data"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"github.com/sirupsen/logrus"
)

// Restaurant представляет структуру данных ресторана.
type Restaurant struct {
	ID       string   `json:"id"`
	Name     string   `json:"name"`
	Address  string   `json:"address"`
	Phone    string   `json:"phone"`
	Location GeoPoint `json:"location"`
}

// GeoPoint представляет географическую точку с долготой и широтой.
type GeoPoint struct {
	Longitude float64 `json:"lon"`
	Latitude  float64 `json:"lat"`
}

// CreateAndMapping создает индекс с маппингом в Elasticsearch.
func CreateAndMapping(es *elasticsearch.Client, indexName string, mapping data.Schema) {
	if _, err := es.Indices.Delete([]string{indexName}); err != nil {
		logrus.Fatalf("Error deleting index: %s\n", err)
	}
	indexSet := data.IndexSettings{Mappings: mapping}
	indexSet.Mappings.Properties.Id.Type = "long"

	mappingBytes, err := json.Marshal(indexSet)
	if err != nil {
		logrus.Printf("Error marshaling JSON: %s\n", err)
	}

	// делаем запрос на создание индекса
	req := esapi.IndicesCreateRequest{
		Index: indexName,
		Body:  strings.NewReader(string(mappingBytes)),
	}

	resp, err := req.Do(context.Background(), es)
	if err != nil {
		logrus.Fatalf("Error creating index: %s\n", err)
	}
	defer resp.Body.Close()

	if resp.IsError() {
		logrus.Fatalf("Error response from server: %s", resp)
	} else {
		logrus.Println("Index creation successful")
	}
}

// CreateBulkRequest создает bulk-запрос для Elasticsearch из данных.
func CreateBulkRequest(lines [][]string) string {
	var bulkRequest strings.Builder
	for i, line := range lines[1:] {
		restaurant := Restaurant{
			ID:      fmt.Sprintf("%d", i+1),
			Name:    line[1],
			Address: line[2],
			Phone:   line[3],
			Location: GeoPoint{
				Longitude: parseFloat(line[4]),
				Latitude:  parseFloat(line[5]),
			},
		}
		meta := []byte(fmt.Sprintf(`{ "index" : { "_index" : "places", "_id" : "%s" } }%s`, restaurant.ID, "\n"))
		data, err := json.Marshal(restaurant)
		if err != nil {
			logrus.Fatalf("Error marshaling document: %s", err)
		}
		bulkRequest.Grow(len(meta) + len(data) + 1)
		bulkRequest.Write(meta)
		bulkRequest.Write(data)
		bulkRequest.WriteByte('\n')
	}

	return bulkRequest.String()
}

// SendBulkRequest отправляет bulk-запрос в Elasticsearch.
func SendBulkRequest(es *elasticsearch.Client, indexName, bulkRequest string) {
	req := esapi.BulkRequest{
		Index: indexName,
		Body:  strings.NewReader(bulkRequest),
	}

	resp, err := req.Do(context.Background(), es)
	if err != nil {
		logrus.Fatalf("Error sending the bulk request: %s\", err")
	}
	defer resp.Body.Close()

	if resp.IsError() {
		logrus.Fatalf("Error response from server: %s", resp)
	} else {
		logrus.Println("Bulk request successful")
	}
}

// parseFloat преобразует строку в float64.
func parseFloat(value string) float64 {
	f, err := strconv.ParseFloat(value, 64)
	if err != nil {
		logrus.Fatalf("Error parsing value: %s\n", err)
	}

	return f
}
