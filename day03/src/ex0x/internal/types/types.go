package types

// Data представляет структуру для хранения данных о местах и информации о страницах.
type Data struct {
	Name     string  `json:"name"`
	Total    int     `json:"total"`
	Places   []Place `json:"places"`
	Page     int     `json:"-"`
	PrevPage int     `json:"prev_page,omitempty"`
	NextPage int     `json:"next_page,omitempty"`
	Last     int     `json:"last_page"`
}

// Place представляет информацию о месте.
type Place struct {
	ID       string   `json:"id"`
	Name     string   `json:"name"`
	Address  string   `json:"address"`
	Phone    string   `json:"phone"`
	Location GeoPoint `json:"location"`
}

// GeoPoint представляет географическую точку с широтой и долготой.
type GeoPoint struct {
	Longitude float64 `json:"lon"`
	Latitude  float64 `json:"lat"`
}

// Store представляет интерфейс для работы с хранилищем данных.
type Store interface {
	GetPlaces(limit, offset int) ([]Place, int, error)
	GetClosestPlaces(lat, lon float64, limit int) ([]Place, error)
}
