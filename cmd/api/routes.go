package main

import (
	"net/http"

	"github.com/rs/cors"
)

func (app *application) routes() http.Handler {
	mux := http.NewServeMux()
	c := cors.New(cors.Options{
		AllowedOrigins: []string{app.config.clientURL},
		AllowedMethods: []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Authorization", "Content-Type"},
	})

	// se van a definir nuestras rutas
	app.authRoutes(mux)
	app.chatRoutes(mux)
	app.wsRoutes(mux)

	return c.Handler(mux)
}

func (app *application) authRoutes(mux *http.ServeMux) {
	mux.HandleFunc("POST /api/auth/register", app.registerUser)
	mux.HandleFunc("POST /api/auth/login", app.loginUser)
}

func (app *application) wsRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/ws", app.ConnectWs)
}

func (app *application) chatRoutes(mux *http.ServeMux) {
	mux.Handle("POST /api/chat/message", app.AuthMiddleware(app.SendMessage))
}
