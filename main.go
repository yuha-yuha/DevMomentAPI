package main

import (
	"log"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/yuha-yuha/DevMomentAPI/controllers/router"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {

	mux := router.Get("./sample.json")

	server := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	server.ListenAndServe()
}
