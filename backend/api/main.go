package main

import (
	"context"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"go_rest.com/rest-ws/api/handlers"
	"go_rest.com/rest-ws/api/server"
	"log"
	"net/http"
	"os"
)

func main() {
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	PORT := os.Getenv("PORT")
	JWT_SECRET := os.Getenv("JWT_SECRET")
	DATABASE_URL := os.Getenv("DATABASE_URL")

	ctx := context.Background()
	conf := &server.Config{
		Port:        PORT,
		JWTSecret:   JWT_SECRET,
		DatabaseUrl: DATABASE_URL,
	}
	s, err := server.NewServer(ctx, conf)

	if err != nil {
		log.Fatal(err)
	}

	s.Start(BindRouters)
}

func BindRouters(s server.Server, r *mux.Router) {
	r.HandleFunc("/", handlers.HomeHandler(s)).Methods(http.MethodGet)
	r.HandleFunc("/health", handlers.HealthHandler(s)).Methods(http.MethodGet, http.MethodPost)
}
