package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"ex0x/internal/types"
	"github.com/sirupsen/logrus"
)

// RecommendHandler возвращает JSON с тремя ближайшими ресторанами.
func RecommendHandler(store types.Store) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Set("Content-Type", "application/json")
		latStr := r.URL.Query().Get("lat")
		lonStr := r.URL.Query().Get("lon")

		lat := parseFloat(latStr)
		lon := parseFloat(lonStr)

		places, err := store.GetClosestPlaces(lat, lon, 3)
		if err != nil {
			http.Error(rw, err.Error(), http.StatusInternalServerError)
			return
		}

		response := struct {
			Name   string        `json:"name"`
			Places []types.Place `json:"places"`
		}{
			Name:   "Recommendation",
			Places: places,
		}

		jsonData, err := json.MarshalIndent(response, "", "  ")
		if err != nil {
			http.Error(rw, err.Error(), http.StatusInternalServerError)
			return
		}

		rw.Write(jsonData)
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
