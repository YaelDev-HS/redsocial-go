package main

import (
	"errors"
	"fmt"
	"net/http"
	"time"

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

	token, err := app.models.Token.Insert(user.ID, time.Hour*24, data.ScopeAuthentication)

	if err != nil {
		app.internalServerError(w, err)
		return
	}

	response := responseBody{
		Data: map[string]any{
			"user":  user,
			"token": token,
		},
	}

	if err := app.writeJson(w, response, http.StatusCreated); err != nil {
		app.internalServerError(w, err)
	}
}

func (app *application) loginUser(w http.ResponseWriter, r *http.Request) {
	var body struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := app.decodeJson(r, &body); err != nil {
		app.badRequest(w, err)
		return
	}

	v := validator.New()
	v.Check(!v.Match(body.Email, validator.EmailRegex), "email", "is not valid")
	v.Check(len(body.Password) > 3, "password", "is too short")
	v.Check(len(body.Password) < 60, "password", "is too long")

	if !v.IsValid() {
		app.badRequest(w, v.Errors())
		return
	}

	user, err := app.models.User.FindByEmail(body.Email)

	if err != nil {
		switch {
		case errors.Is(err, data.ErrModelNotFound):
			app.notFound(w, fmt.Errorf("email or password is not valid"))
		default:
			app.internalServerError(w, err)
		}

		return
	}

	ok, err := user.ComparePassword(body.Password)

	if err != nil {
		app.internalServerError(w, err)
		return
	}

	if !ok {
		app.badRequest(w, fmt.Errorf("email or password is not valid"))
		return
	}

	token, err := app.models.Token.Insert(user.ID, time.Hour*48, data.ScopeAuthentication)

	if err != nil {
		app.internalServerError(w, err)
		return
	}

	response := responseBody{
		Data: map[string]any{
			"user":  user,
			"token": token,
		},
	}

	if err := app.writeJson(w, response, http.StatusOK); err != nil {
		app.internalServerError(w, err)
		return
	}
}

func (app *application) checkToken(w http.ResponseWriter, r *http.Request) {
	app.writeJson(w, responseBody{
		Data: map[string]any{
			"ok": true,
		},
	}, http.StatusOK)
}
