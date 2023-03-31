package main

import (
	"net/http"

	"github.com/justinas/alice"
)

func (app *application) routes() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/snippet", app.showSnippet)
	mux.HandleFunc("/snippet/create", app.createSnippet)

	fileserver := http.FileServer(http.Dir("ui/static"))
	mux.Handle("/static/", http.StripPrefix("/static", fileserver))
	standardMiddleware := alice.New(app.recoverPanic, app.logRequest, secureHeader)
	return standardMiddleware.Then(mux)
}
