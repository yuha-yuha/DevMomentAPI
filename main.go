package main

import (
	"fmt"
	"net/http"

	"github.com/yuha-yuha/DevMomentAPI/controllers/router"
)

func main() {

	mux := router.Get("./sample.json")

	server := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	fmt.Print("Running Server!!\n\n")
	server.ListenAndServe()
}
