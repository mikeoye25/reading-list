package main

import (
	"net/http"

	"reading-list/internal/data"
)

func (app *application) registerUserHandler(w http.ResponseWriter, r *http.Request) {

	
	if r.Method == http.MethodPost {
		var input struct {
			Name     string `json:"name"`
			Username string `json:"username"`
			Email    string `json:"email"`
			Password string `json:"password"`
		}

		err := app.readJSON(w, r, &input)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}

		user := &data.User{
			Name:     input.Name,
			Username: input.Username,
			Email:     input.Email,
			Password:    input.Password,
		}

		err = app.models.Users.Insert(user)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		headers := make(http.Header)
		// headers.Set("Location", fmt.Sprintf("v1/books/%d", book.ID))

		// Write the JSON response with a 201 Created status code and the Location header set.
		err = app.writeJSON(w, http.StatusCreated, envelope{"user": user}, headers)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
	}
}
