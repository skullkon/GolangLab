package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/skullkon/go_lab/internal/app/api"
)

var (
	configPath string
)

func main() {
	err := godotenv.Load()
	config := api.NewConfig()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	server := api.New(config)
	log.Println("Starting")
	log.Fatal(server.Start())
}
