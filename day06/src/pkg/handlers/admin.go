package handlers

import (
	"html/template"
	"log"
	"net/http"

	data "fortress-of-solitude/pkg/db"

	"github.com/jackc/pgx/v5"
)

const adminHTML = "./templates/admin.html"

func AdminPage(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles(adminHTML))
	tmpl.Execute(w, nil)
}

func CreatePost(w http.ResponseWriter, r *http.Request, db *pgx.Conn) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	title := r.FormValue("title")
	content := r.FormValue("content")
	err := data.CreatePost(db, title, content)
	if err != nil {
		http.Error(w, "Failed to create post", http.StatusInternalServerError)
		log.Println(err)
		return
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
