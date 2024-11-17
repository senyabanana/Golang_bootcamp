package main

import (
	"net/http"

	"ex0x/internal/db"
	"ex0x/internal/handlers"
	"ex0x/internal/middleware"
	"ex0x/internal/types"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/sirupsen/logrus"
)

var store types.Store

func main() {
	es, err := elasticsearch.NewDefaultClient()
	if err != nil {
		logrus.Fatalf("Error creating the client: %s", err)
	}

	err = db.SetMaxResultWindow(es, "places", 20000)
	if err != nil {
		logrus.Fatalf("Error setting max_result_window: %s", err)
	}

	store = db.NewElasticStore(es)

	http.HandleFunc("/", handlers.PlacesHandler(store))
	http.HandleFunc("/api/places", handlers.APIPlacesHandler(store))
	http.HandleFunc("/api/get_token", handlers.GetTokenHandler(store))
	//http.HandleFunc("/api/recommend", handlers.RecommendHandler(store)) // <-- ex03

	// Применение JWTMiddleware для защиты конечной точки /api/recommend.
	recommendHandler := middleware.JWTMiddleware(http.HandlerFunc(handlers.RecommendHandler(store)))
	http.Handle("/api/recommend", recommendHandler)
	http.ListenAndServe(`:8888`, nil)
}
