package main

import (
	"html/template"
	"net/http"
)

func (app *application) spotik(w http.ResponseWriter, r *http.Request) {
	app.render(w, "./ui/html/helpyouout/home.html")
}

func (app *application) render(w http.ResponseWriter, pagePath string) {
	files := []string{
		pagePath,
		"./ui/html/base.html",
		"./ui/html/partials/navbar.html",
		"./ui/html/partials/music-player.html",
		"./ui/static/style.css",
	}
	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.serverError(w, err)
		return
	}
	err = ts.Execute(w, nil)
	if err != nil {
		app.serverError(w, err)
	}
}
