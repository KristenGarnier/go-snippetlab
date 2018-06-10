package main

import (
	"net/http"
)

func (app *App) Routes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", app.Home)
	mux.HandleFunc("/snippet", app.ShowSnippet)
	mux.HandleFunc("/snippet/new", app.NewSnippet)

	httpFileServer := http.FileServer(http.Dir(app.StaticDir))
	mux.Handle("/static/", http.StripPrefix("/static", httpFileServer))

	return mux
}
