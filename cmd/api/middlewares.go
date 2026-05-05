package main

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/YaelDev-HS/redsocial-go/internal/data"
)

func (app *application) AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		header := r.Header.Get("Authorization")

		if header == "" || !strings.HasPrefix(header, "Bearer ") {
			app.unauthorized(w, fmt.Errorf("missing token"))
			return
		}

		plainText := header[7:]

		if len(plainText) != 26 {
			app.unauthorized(w, fmt.Errorf("token is not valid or expired"))
			return
		}

		token, err := app.models.Token.FindByPlaintext(plainText, data.ScopeAuthentication)

		if err != nil {
			switch {
			case errors.Is(err, data.ErrModelNotFound):
				app.badRequest(w, fmt.Errorf("token is not valid or expired"))
			default:
				app.internalServerError(w, err)
			}

			return
		}

		//TODO: guardar el usuario en el contexto de Go
		fmt.Println(token.User)

		next.ServeHTTP(w, r)
	}
}
