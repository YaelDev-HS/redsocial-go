package main

import (
	"errors"
	"net/http"

	"github.com/YaelDev-HS/redsocial-go/internal/data"
	"github.com/YaelDev-HS/redsocial-go/internal/validator"
)

func (app *application) registerUser(w http.ResponseWriter, r *http.Request) {
	var body struct {
		Email    string `json:"email"`
		Password string `json:"password"`
		Username string `json:"username"`
	}

	err := app.decodeJson(r, &body)

	if err != nil {
		app.badRequest(w, err)
		return
	}

	v := validator.New()

	v.Check(body.Email != "", "email", "email is not valid")
	v.Check(body.Password != "", "password", "password is not valid")
	v.Check(body.Username != "", "username", "username is not valid")

	if !v.IsValid() {
		app.badRequest(w, v.Errors())
		return
	}

	user := &data.User{
		Username:  body.Username,
		Email:     body.Email,
		Nickname:  body.Username,
		IsEnabled: true,
	}

	err = user.SetPassword(body.Password)

	if err != nil {
		app.internalServerError(w, err)
		return
	}

	err = app.models.User.Create(user)

	if err != nil {
		switch {
		case errors.Is(err, data.ErrDuplicatedKey):
			app.badRequest(w, "email or username already exists")
		default:
			app.internalServerError(w, err)
		}

		return
	}

	response := responseBody{
		Data: map[string]any{
			"user": user,
		},
	}

	if err := app.writeJson(w, response, http.StatusCreated); err != nil {
		app.internalServerError(w, err)
	}
}
