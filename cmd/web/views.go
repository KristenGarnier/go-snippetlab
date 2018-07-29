package main

import (
	"bytes"
	"html/template"
	"net/http"
	"path/filepath"
	"time"

	"go-snippetlab/pkg/models"
)

type HTMLData struct {
	Flash    string
	Form     interface{}
	Path     string
	Snippet  *models.Snippet
	Snippets []*models.Snippet
}

func (app *App) RenderHTML(w http.ResponseWriter, r *http.Request, page string, data *HTMLData) {
	if data == nil {
		data = &HTMLData{}
	}

	data.Path = r.URL.Path

	files := []string{
		filepath.Join(app.HTMLDir, "base.html"),
		filepath.Join(app.HTMLDir, page),
	}

	fm := template.FuncMap{
		"humanDate": humanDate,
	}

	ts, err := template.New("").Funcs(fm).ParseFiles(files...)
	if err != nil {
		app.ServerError(w, err)
		return
	}

	buf := new(bytes.Buffer)

	err = ts.ExecuteTemplate(buf, "base", data)
	if err != nil {
		app.ServerError(w, err)
		return
	}

	buf.WriteTo(w)
}

func humanDate(t time.Time) string {
	return t.Format("02 Jan 2006 at 15:04")
}
