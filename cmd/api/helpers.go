package main

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/YaelDev-HS/redsocial-go/internal/data"
)

var (
	unauthorizedError = errors.New("token is not valid or expired")
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

func (app *application) GetParam(r *http.Request, key string) (string, bool) {
	value := r.URL.Query().Get(key)

	if value == "" {
		return value, false
	}

	return value, true
}

func (app *application) getUserAuthByToken(plaintext string) (*data.User, error) {
	if len(plaintext) != 26 {
		return nil, unauthorizedError
	}

	token, err := app.models.Token.FindByPlaintext(plaintext, data.ScopeAuthentication)

	if err != nil {
		if errors.Is(err, data.ErrModelNotFound) {
			return nil, unauthorizedError
		}

		return nil, err
	}

	return token.User, nil
}
