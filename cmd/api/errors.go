package main

import (
	"fmt"
	"net/http"
)

func (app *application) unauthorized(w http.ResponseWriter, err error) {
	app.httpError(w, http.StatusUnauthorized, err.Error())
}

func (app *application) badRequest(w http.ResponseWriter, err any) {
	app.httpError(w, http.StatusBadRequest, err)
}

func (app *application) notFound(w http.ResponseWriter, err error) {
	app.httpError(w, http.StatusNotFound, err.Error())
}

func (app *application) internalServerError(w http.ResponseWriter, err error) {
	//TODO: logger
	fmt.Printf("internal error: %s\n", err)
	app.httpError(w, http.StatusInternalServerError, "internal server error!")
}

func (app *application) httpError(w http.ResponseWriter, status int, err any) {
	response := responseBody{
		Error: true,
		Data: map[string]any{
			"error": err,
		},
	}

	app.writeJson(w, response, status)
}
