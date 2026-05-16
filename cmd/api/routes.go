package main

import "net/http"

func (app *application) routes() *http.ServeMux {
	mux := http.NewServeMux()

	// se van a definir nuestras rutas
	app.authRoutes(mux)
	app.chatRoutes(mux)
	app.wsRoutes(mux)

	return mux
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
