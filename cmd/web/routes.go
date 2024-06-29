package main

import (
	"net/http"

	"github.com/justinas/alice"
)

func (app *application) routes() http.Handler {
	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("GET /static/", http.StripPrefix("/static/", fileServer))

	dynamic := alice.New(app.sessionManager.LoadAndSave, app.authenticate)
	protected := dynamic.Append(app.requireAuthentication)

	// snippet
	mux.Handle("GET /", dynamic.ThenFunc(app.homeHandler))
	mux.Handle("GET /snippet/view/{id}", dynamic.ThenFunc(app.snippetViewHandler))
	mux.Handle("GET /snippet/create", protected.ThenFunc(app.snippetCreateHandler))
	mux.Handle("POST /snippet/create", protected.ThenFunc(app.snippetCreatePostHandler))
	// user
	mux.Handle("GET /user/signup", dynamic.ThenFunc(app.userSignup))
	mux.Handle("POST /user/signup", dynamic.ThenFunc(app.userSignupPost))
	mux.Handle("POST /user/login", dynamic.ThenFunc(app.userLoginPost))
	mux.Handle("GET /user/login", dynamic.ThenFunc(app.userLogin))
	mux.Handle("POST /user/logout", protected.ThenFunc(app.userLogoutPost))

	standard := alice.New(app.recoverPanic, app.logRequest, secureHeaders)

	return standard.Then(mux)
}
