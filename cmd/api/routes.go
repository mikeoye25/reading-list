package main

import "net/http"

func (app *application) routes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", http.NotFound) // Catch-all route
	mux.HandleFunc("/v1/healthcheck", app.healthcheck)
	// mux.HandleFunc("/v1/books", app.getCreateBooksHandler)
	// mux.HandleFunc("/v1/books", middleware(app.getCreateBooksHandler))
	mux.HandleFunc("/v1/books", TokenVerifyMiddleware(app.getCreateBooksHandler))
	mux.HandleFunc("/v1/books/", app.getUpdateDeleteBooksHandler)

	mux.HandleFunc("/v1/auth/register", app.registerUserHandler)
	mux.HandleFunc("/v1/auth/generate", app.generateTokenHandler)

	return mux
}