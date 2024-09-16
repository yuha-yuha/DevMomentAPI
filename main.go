package main

import (
	"net/http"

	"github.com/yuha-yuha/DevMomentAPI/controllers/router"
)

func main() {

	mux := router.Get("./sample.json")

	server := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	server.ListenAndServe()
}
