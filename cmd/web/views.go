package main

import (
	"html/template"
	"net/http"
	"path/filepath"

	"snippetbox.org/pkg/models"
)

type HTMLData struct {
	Snippet *models.Snippet
}

func (app *App) RenderHTML(w http.ResponseWriter, page string, data *HTMLData) {
	files := []string{
		filepath.Join(app.HTMLDir, "base.html"),
		filepath.Join(app.HTMLDir, page),
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.ServerError(w, err)
		return
	}

	err = ts.ExecuteTemplate(w, "base", data)
	if err != nil {
		app.ServerError(w, err)
		return
	}
}
