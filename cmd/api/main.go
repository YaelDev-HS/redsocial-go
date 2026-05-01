package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/YaelDev-HS/redsocial-go/internal/data"
	"github.com/joho/godotenv"
)

type application struct {
	models *data.Models
}

func init() {
	godotenv.Load()
}

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	db, err := ConnectDB(os.Getenv("DB_DSN"))

	if err != nil {
		log.Fatalf("internal error (DB) = %s\n", err)
		return
	}

	models := data.New(db)

	app := &application{
		models: models,
	}

	server := http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: app.routes(),
	}

	log.Printf("Server running on port: %s\n", port)

	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("internal server error: %s\n", err)
	}
}
