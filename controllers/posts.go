package controllers

import (
	"net/http"
	"text/template"
)

func Home(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("view/index.html")

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	err = t.Execute(w, nil)
}

func Allpost(w http.ResponseWriter, r *http.Request) {
	var n Post
	// db.Client.Database()
	n.New()
}

func AddPost() {}
