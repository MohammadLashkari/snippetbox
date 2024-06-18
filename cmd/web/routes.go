package main

import "net/http"

func (app *application) routes() *http.ServeMux {
	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static/", fileServer))

	mux.HandleFunc("/", app.homeHandler)
	mux.HandleFunc("/snippet/view", app.snippetViewHandler)
	mux.HandleFunc("POST /snippet/create", app.snippetCreateHandler)

	return mux
}
