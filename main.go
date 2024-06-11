package main

import (
	"log"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	w.Write([]byte("hello from snippetbox"))
}

func snippetViewHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("display a specific snippet..."))
}

func snippetCreateHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("create a new snippet..."))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /", homeHandler)
	mux.HandleFunc("GET /snippet/view", snippetViewHandler)
	mux.HandleFunc("GET /snippet/hello/", snippetViewHandler)
	mux.HandleFunc("GET /snippet/create", snippetCreateHandler)

	server := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
	log.Println("starign server on :8080")
	log.Fatal(server.ListenAndServe())
}
