package main

import (
	"net/http"

	"github.com/yuha-yuha/DevMomentAPI/controller/router"
)

func main() {

	mux := router.Get()

	server := http.Server{
		Addr:    ":8081",
		Handler: mux,
	}

	server.ListenAndServe()
}
