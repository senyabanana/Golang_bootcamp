package handlers

import (
	"encoding/json"
	"fmt"
	"math"
	"net/http"
	"strconv"

	"ex0x/internal/types"
)

// APIPlacesHandler возвращает HTTP-обработчик для API, который возвращает список мест в формате JSON.
func APIPlacesHandler(store types.Store) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Set("Content-Type", "application/json")
		var (
			res types.Data
			err error
		)

		pageStr := r.URL.Query().Get("page")
		res.Page, err = strconv.Atoi(pageStr)
		if err != nil || res.Page < 1 {
			http.Error(rw, fmt.Sprintf("Invalid 'page' value: '%s'", pageStr), http.StatusBadRequest)
			return
		}

		limit := 10
		offset := (res.Page - 1) * limit

		res.Places, res.Total, err = store.GetPlaces(limit, offset)
		if err != nil {
			http.Error(rw, err.Error(), http.StatusInternalServerError)
			return
		}
		res.Last = int(math.Ceil(float64(res.Total) / float64(limit)))

		if res.Page > res.Last {
			http.Error(rw, fmt.Sprintf("Invalid 'page' value: '%s'", pageStr), http.StatusBadRequest)
			return
		}

		if res.Page > 1 {
			res.PrevPage = res.Page - 1
		}
		if res.Page < res.Last {
			res.NextPage = res.Page + 1
		}

		res.Name = "Places"

		jsonData, err := json.MarshalIndent(res, "", "  ")
		if err != nil {
			http.Error(rw, err.Error(), http.StatusInternalServerError)
			return
		}

		rw.Write(jsonData)
	}
}
