package main

import (
	"net/http"

	"github.com/YaelDev-HS/redsocial-go/internal/data"
	"github.com/YaelDev-HS/redsocial-go/internal/store"
	"github.com/YaelDev-HS/redsocial-go/internal/validator"
)

func (app *application) SendMessage(w http.ResponseWriter, r *http.Request) {
	var body struct {
		Message string `json:"message"`
	}

	if err := app.decodeJson(r, &body); err != nil {
		app.badRequest(w, err.Error())
		return
	}

	v := validator.New()
	v.Check(body.Message != "", "message", "missing value")

	if !v.IsValid() {
		app.badRequest(w, v.Errors())
		return
	}

	user, _ := app.getUserContext(r)

	message := &data.Message{
		Message:   body.Message,
		UserID:    user.ID,
		IsEnabled: true,
	}

	err := app.models.Message.Create(message)

	if err != nil {
		app.internalServerError(w, err)
		return
	}

	message.User = &data.User{
		ID:       user.ID,
		Nickname: user.Nickname,
	}

	//? notificar mediante websockets
	storeMsg := &store.StoreMessage{
		Type:    store.NEW_MESSAGE,
		Message: message,
	}
	app.store.NotifyAll(storeMsg)

	response := responseBody{
		Data: map[string]any{
			"user":    user,
			"message": message,
		},
	}

	app.writeJson(w, response, http.StatusCreated)
}
