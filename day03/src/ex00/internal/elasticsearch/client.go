package elasticsearch

import (
	"github.com/elastic/go-elasticsearch/v8"
)

// NewClient создает и возвращает новый клиент Elasticsearch.
func NewClient() (*elasticsearch.Client, error) {
	es, err := elasticsearch.NewDefaultClient()
	if err != nil {
		return nil, err
	}
	return es, nil
}
