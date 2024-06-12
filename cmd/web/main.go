package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", homeHandler)
	mux.HandleFunc("/snippet/view", snippetViewHandler)
	mux.HandleFunc("/snippet/create", snippetCreateHandler)

	server := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
	log.Println("startig server on :8080")
	log.Fatal(server.ListenAndServe())
}
