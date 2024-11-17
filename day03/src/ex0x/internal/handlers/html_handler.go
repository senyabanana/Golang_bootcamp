package handlers

import (
	"fmt"
	"html/template"
	"math"
	"net/http"
	"strconv"

	"ex0x/internal/types"
)

const placesHTML = `template/places.html`

// PlacesHandler возвращает HTTP-обработчик для HTML, который возвращает список мест в формате HTML.
func PlacesHandler(store types.Store) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
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

		tmpl, err := template.New("places.html").Funcs(
			template.FuncMap{
				"sum": sum,
				"sub": sub,
			},
		).ParseFiles(placesHTML)

		if err != nil {
			http.Error(rw, err.Error(), http.StatusInternalServerError)
			return
		}

		res.Name = "Places"

		if err := tmpl.Execute(rw, res); err != nil {
			http.Error(rw, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

// sum возвращает сумму двух чисел.
func sum(x, y int) int {
	return x + y
}

// sub возвращает разницу между двумя числами.
func sub(x, y int) int {
	return x - y
}
