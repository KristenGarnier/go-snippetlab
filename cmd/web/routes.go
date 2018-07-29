package main

import (
	"net/http"

	"github.com/bmizerany/pat"
)

func (app *App) Routes() *pat.PatternServeMux {
	mux := pat.New()
	mux.Get("/", http.HandlerFunc(app.Home))
	mux.Get("/snippet/new", http.HandlerFunc(app.NewSnippet))
	mux.Post("/snippet/new", http.HandlerFunc(app.CreateSnippet))
	mux.Get("/snippet/:id", http.HandlerFunc(app.ShowSnippet))

	httpFileServer := http.FileServer(http.Dir(app.StaticDir))
	mux.Get("/static/", http.StripPrefix("/static", httpFileServer))

	return mux
}
