package main

import "net/http"

func (app *application) routes() *http.ServeMux {
	mux := http.NewServeMux()

	// se van a definir nuestras rutas
	app.authRoutes(mux)

	return mux
}

func (app *application) authRoutes(mux *http.ServeMux) {
	mux.HandleFunc("POST /api/auth/register", app.registerUser)
}
