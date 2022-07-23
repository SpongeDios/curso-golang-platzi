package main

import (
	"05-rest/handlers"
	"05-rest/server"
	"context"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port := os.Getenv("PORT")
	jwtSecret := os.Getenv("JWT_SECRET")
	databaseUrl := os.Getenv("DATABASE_URL")

	server, err := server.NewServer(context.Background(), &server.Config{
		Port:        port,
		JWTSecret:   jwtSecret,
		DatabaseUrl: databaseUrl,
	})

	if err != nil {
		log.Fatal(err)
	}
	server.Start(BindRoutes)
}

func BindRoutes(s server.Server, r *mux.Router){
	r.HandleFunc("/", handlers.HomeHandle(s))
}
