package main

import (
	"log"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from SnippetBox"))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", Home)

	log.Println("Starting Server on port :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
