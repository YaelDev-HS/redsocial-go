package main

import (
	"encoding/json"
	"net/http"
)

type responseBody struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

func (app *application) writeJson(w http.ResponseWriter, body responseBody, status int) error {
	value, err := json.Marshal(body)

	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(value)

	return nil
}

func (app *application) decodeJson(r *http.Request, body any) error {
	err := json.NewDecoder(r.Body).Decode(body)

	if err != nil {
		//TODO: evaluar error
		return err
	}

	return nil
}
