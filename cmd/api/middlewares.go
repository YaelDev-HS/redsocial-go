package main

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
)

func (app *application) AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		header := r.Header.Get("Authorization")

		if header == "" || !strings.HasPrefix(header, "Bearer ") {
			app.unauthorized(w, fmt.Errorf("missing token"))
			return
		}

		plainText := header[7:]
		user, err := app.getUserAuthByToken(plainText)

		if err != nil {
			switch {
			case errors.Is(err, unauthorizedError):
				app.unauthorized(w, err)
			default:
				app.internalServerError(w, err)
			}

			return
		}

		request := app.setUserContext(r, user)
		next.ServeHTTP(w, request)
	}
}
