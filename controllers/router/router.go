package router

import (
	"net/http"

	"github.com/yuha-yuha/DevMomentAPI/controllers/handlers"
	"github.com/yuha-yuha/DevMomentAPI/controllers/middlewares"
	"github.com/yuha-yuha/DevMomentAPI/services"
)

func Get(filePath string) *http.ServeMux {
	services.SubscribeUserDefineAPIs(filePath)
	mux := http.NewServeMux()
	udHandlers := handlers.CreateUserDefineHandler(services.GetUserDefineAPIMap(), services.GetUserDefineModels(filePath))

	for _, udh := range udHandlers {
		mux.Handle(udh.Path, middlewares.AccessLogger(udh.HandlerFunc))
	}

	mux.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		s := "hello"
		w.Write([]byte(s))
	})
	return mux
}
