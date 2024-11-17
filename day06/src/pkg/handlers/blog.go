package handlers

import (
	"html/template"
	"log"
	"net/http"

	dataDb "fortress-of-solitude/pkg/db"

	"github.com/jackc/pgx/v5"
	"github.com/russross/blackfriday"
)

const (
	indexHTML = "./templates/index.html"
	postHTML  = "./templates/post.html"
)

func ShowPosts(w http.ResponseWriter, r *http.Request, db *pgx.Conn) {
	posts, err := dataDb.GetPosts(db, 3, 0)
	if err != nil {
		http.Error(w, "Unable to fetch posts", http.StatusInternalServerError)
		log.Println(err)
		return
	}

	tmpl := template.Must(template.ParseFiles(indexHTML))
	tmpl.Execute(w, posts)
}

func ShowSinglePost(w http.ResponseWriter, r *http.Request, db *pgx.Conn) {
	id := r.URL.Query().Get("id")
	post, err := dataDb.GetPostByID(db, id)
	if err != nil {
		http.Error(w, "Post not found", http.StatusNotFound)
		log.Println(err)
		return
	}

	content := blackfriday.MarkdownCommon([]byte(post.Content))
	data := struct {
		Title   string
		Content template.HTML
	}{
		Title:   post.Title,
		Content: template.HTML(content),
	}

	tmpl := template.Must(template.ParseFiles(postHTML))
	tmpl.Execute(w, data)
}
