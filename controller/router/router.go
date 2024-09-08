package router

import (
	"net/http"

	"github.com/yuha-yuha/DevMomentAPI/controller/handler"
	"github.com/yuha-yuha/DevMomentAPI/lib"
)

func Get() *http.ServeMux {
	mux := http.NewServeMux()
	udHandlers := handler.CreateUserDefineHandler(lib.GetUserDefineAPIs(), lib.GetUserDefineModels())
	//udHandlers2 := handler.GetUserDefineHandlers(lib.JsonParse())
	for _, udh := range udHandlers {
		mux.Handle(udh.Path, udh.HandlerFunc)
	}

	mux.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		s := "hello"
		w.Write([]byte(s))
	})
	return mux
}
