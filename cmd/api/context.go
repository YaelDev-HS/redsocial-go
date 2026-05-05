package main

import (
	"context"
	"net/http"

	"github.com/YaelDev-HS/redsocial-go/internal/data"
)

type KeyContext string

var (
	KeyUserCtx KeyContext = "user_ctx"
)

func (app *application) setContext(r *http.Request, value any, key KeyContext) *http.Request {
	ctx := context.WithValue(r.Context(), key, value)
	return r.WithContext(ctx)
}

func (app *application) setUserContext(r *http.Request, user *data.User) *http.Request {
	return app.setContext(r, user, KeyUserCtx)
}

func (app *application) getUserContext(r *http.Request) (*data.User, bool) {
	value := r.Context().Value(KeyUserCtx)

	if value == nil {
		return nil, false
	}

	user, ok := value.(*data.User)

	if !ok {
		return nil, false
	}

	return user, true
}
