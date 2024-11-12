package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/yuha-yuha/DevMomentAPI/controllers/router"
)

func main() {

	flag.Parse()
	mux := router.Get("./" + flag.Arg(0))

	server := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	fmt.Println("Running Server!!")
	fmt.Print("source: \"" + flag.Arg(0) + "\"\n\n")
	server.ListenAndServe()
}
