package main

import "net/http"

func (app *application) routes() http.Handler {
	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("GET /static/", http.StripPrefix("/static/", fileServer))

	mux.HandleFunc("GET /", app.homeHandler)
	mux.HandleFunc("GET /snippet/view/{id}", app.snippetViewHandler)
	mux.HandleFunc("GET /snippet/create", app.snippetCreateHandler)
	mux.HandleFunc("POST /snippet/create", app.snippetCreatePostHandler)

	stack := middlewareStack(
		app.recoverPanic,
		app.logRequest,
		secureHeaders,
	)
	return stack(mux)
}
