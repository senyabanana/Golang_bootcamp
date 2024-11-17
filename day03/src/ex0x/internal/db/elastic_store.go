package db

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"ex0x/internal/types"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
)

// ElasticStore представляет хранилище данных с использованием Elasticsearch.
type ElasticStore struct {
	client *elasticsearch.Client
}

// NewElasticStore создает новый экземпляр ElasticStore.
func NewElasticStore(client *elasticsearch.Client) *ElasticStore {
	return &ElasticStore{client: client}
}

// GetPlaces возвращает список мест из Elasticsearch с заданным лимитом и смещением.
func (es *ElasticStore) GetPlaces(limit int, offset int) ([]types.Place, int, error) {
	total, err := es.getTotalPlaces()
	if err != nil {
		return nil, 0, err
	}

	query := fmt.Sprintf(`{
        "from": %d,
        "size": %d,
        "sort": [
            {"id": {"order": "asc"}}
        ]
    }`, offset, limit)

	res, err := es.client.Search(
		es.client.Search.WithContext(context.Background()),
		es.client.Search.WithIndex("places"),
		es.client.Search.WithBody(strings.NewReader(query)),
		es.client.Search.WithTrackTotalHits(true),
	)
	if err != nil {
		return nil, 0, err
	}
	defer res.Body.Close()

	if res.IsError() {
		return nil, 0, err
	}

	var r struct {
		Hits struct {
			Total struct {
				Value int `json:"value"`
			} `json:"total"`
			Hits []struct {
				Source types.Place `json:"_source"`
			} `json:"hits"`
		} `json:"hits"`
	}

	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		return nil, 0, err
	}

	places := make([]types.Place, len(r.Hits.Hits))
	for i, hit := range r.Hits.Hits {
		places[i] = hit.Source
	}

	return places, total, nil
}

// getTotalPlaces возвращает общее количество мест в индексе Elasticsearch.
func (es *ElasticStore) getTotalPlaces() (int, error) {
	res, err := es.client.Count(
		es.client.Count.WithContext(context.Background()),
		es.client.Count.WithIndex("places"),
	)
	if err != nil {
		return 0, err
	}
	defer res.Body.Close()

	if res.IsError() {
		return 0, err
	}

	var r struct {
		Count int `json:"count"`
	}

	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		return 0, err
	}

	return r.Count, nil
}

// GetClosestPlaces возвращает три ближайших ресторана к заданным координатам.
func (es *ElasticStore) GetClosestPlaces(lat, lon float64, limit int) ([]types.Place, error) {
	indexName := "places"
	query := map[string]interface{}{
		"size": limit,
		"query": map[string]interface{}{
			"match_all": map[string]interface{}{},
		},
		"sort": []map[string]interface{}{
			{
				"_geo_distance": map[string]interface{}{
					"location": map[string]float64{
						"lat": lat,
						"lon": lon,
					},
					"order":           "asc",
					"unit":            "km",
					"mode":            "min",
					"distance_type":   "arc",
					"ignore_unmapped": true,
				},
			},
		},
	}

	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		return nil, err
	}

	res, err := es.client.Search(
		es.client.Search.WithContext(context.Background()),
		es.client.Search.WithIndex(indexName),
		es.client.Search.WithBody(&buf),
		es.client.Search.WithTrackTotalHits(true),
	)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.IsError() {
		return nil, err
	}

	var r struct {
		Hits struct {
			Hits []struct {
				Source types.Place `json:"_source"`
			} `json:"hits"`
		} `json:"hits"`
	}

	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		return nil, err
	}

	places := make([]types.Place, len(r.Hits.Hits))
	for i, hit := range r.Hits.Hits {
		places[i] = hit.Source
	}

	return places, nil
}

// SetMaxResultWindow устанавливает максимальное количество возвращаемых результатов для указанного индекса.
func SetMaxResultWindow(client *elasticsearch.Client, indexName string, maxResultWindow int) error {
	settings := fmt.Sprintf(`{
        "index": {
            "max_result_window": %d
        }
    }`, maxResultWindow)

	req := esapi.IndicesPutSettingsRequest{
		Index: []string{indexName},
		Body:  strings.NewReader(settings),
	}

	res, err := req.Do(context.Background(), client)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.IsError() {
		return err
	}

	return nil
}
