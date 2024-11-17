package db

import (
	"context"
	"log"
	"time"

	"github.com/jackc/pgx/v5"
)

type Post struct {
	ID        int
	Title     string
	Content   string
	CreatedAt time.Time
}

func InitDB(conn *pgx.Conn) {
	query := `
    CREATE TABLE IF NOT EXISTS articles (
        id SERIAL PRIMARY KEY,
        title TEXT NOT NULL,
        content TEXT NOT NULL,
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
    );`
	_, err := conn.Exec(context.Background(), query)
	if err != nil {
		log.Fatalf("Failed to create table: %v", err)
	}
}

func GetPosts(db *pgx.Conn, limit, offset int) ([]Post, error) {
	rows, err := db.Query(context.Background(), "SELECT id, title, content, created_at FROM articles ORDER BY created_at DESC LIMIT $1 OFFSET $2", limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []Post
	for rows.Next() {
		var post Post
		if err := rows.Scan(&post.ID, &post.Title, &post.Content, &post.CreatedAt); err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}
	return posts, nil
}

func GetPostByID(db *pgx.Conn, id string) (Post, error) {
	var post Post
	query := "SELECT id, title, content, created_at FROM articles WHERE id = $1"
	err := db.QueryRow(context.Background(), query, id).Scan(&post.ID, &post.Title, &post.Content, &post.CreatedAt)
	if err != nil {
		return Post{}, err
	}
	return post, nil
}

func CreatePost(db *pgx.Conn, title, content string) error {
	_, err := db.Exec(context.Background(), "INSERT INTO articles (title, content, created_at) VALUES ($1, $2, NOW())", title, content)
	return err
}
