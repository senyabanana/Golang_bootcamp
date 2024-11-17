package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"fortress-of-solitude/pkg/db"
	"fortress-of-solitude/pkg/handlers"

	"github.com/jackc/pgx/v5"
	"golang.org/x/time/rate"
)

func main() {
	dbURL := "postgres://user:pass@localhost:5432/articles"

	conn, err := pgx.Connect(context.Background(), dbURL)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}
	defer conn.Close(context.Background())

	db.InitDB(conn)

	var limiter = rate.NewLimiter(rate.Limit(100), 50)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if limiter.Allow() {
			handlers.ShowPosts(w, r, conn)
		} else {
			http.Error(w, "Too Many Requests", http.StatusTooManyRequests)
		}
	})
	http.HandleFunc("/post", func(w http.ResponseWriter, r *http.Request) {
		if limiter.Allow() {
			handlers.ShowSinglePost(w, r, conn)
		} else {
			http.Error(w, "Too Many Requests", http.StatusTooManyRequests)
		}
	})
	http.HandleFunc("/admin", handlers.AdminPage)
	http.HandleFunc("/admin/submit", func(w http.ResponseWriter, r *http.Request) {
		if limiter.Allow() {
			handlers.CreatePost(w, r, conn)
		} else {
			http.Error(w, "Too Many Requests", http.StatusTooManyRequests)
		}
	})

	// Запуск сервера
	fmt.Println("Server is running on http://localhost:8888")
	log.Fatal(http.ListenAndServe(":8888", nil))
}
