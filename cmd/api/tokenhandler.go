package main

import (
	"errors"
	"net/http"
)

type TokenRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (app *application) generateTokenHandler(w http.ResponseWriter, r *http.Request) {
	
	var input TokenRequest

	err := app.readJSON(w, r, &input)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	user, err := app.models.Users.Get(input.Email)
	if err != nil {
		switch {
		case errors.Is(err, errors.New("record not found")):
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		default:
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
		return
	}

	err = user.CheckPassword(input.Password)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	tokenString, err:= GenerateJWT(user.Email, user.Username)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	// app.logger.Printf("database connection pool established 5")
	if err := app.writeJSON(w, http.StatusOK, envelope{"token": tokenString}, nil); err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}
