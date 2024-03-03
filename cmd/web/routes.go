package main

import (
	"github.com/bmizerany/pat"
	"github.com/justinas/alice"
	"net/http"
)

func (app *application) routes() http.Handler {
	standardMiddleware := alice.New(app.recoverPanic, app.logRequest, secureHeaders)
	mux1 := pat.New()
	mux1.Get("/", http.HandlerFunc(app.spotik))

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux1.Get("/static/", http.StripPrefix("/static", fileServer))

	return standardMiddleware.Then(mux1)

}
