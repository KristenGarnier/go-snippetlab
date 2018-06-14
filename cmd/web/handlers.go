package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func (app *App) Home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	app.RenderHTML(w, "home.page.html")
}

func (app *App) ShowSnippet(w http.ResponseWriter, r *http.Request) {

	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		app.NotFound(w)
		return
	}

	snippet, err := app.Database.GetSnippet(id)
	if err != nil {
		app.NotFound(w)
		return
	}

	fmt.Fprint(w, snippet)
}

func (app *App) NewSnippet(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display the new snippet form"))
}
