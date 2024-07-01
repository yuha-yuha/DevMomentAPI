package router

import (
	"net/http"

	"github.com/yuha-yuha/DevMomentAPI/controller/handler"
	"github.com/yuha-yuha/DevMomentAPI/helper"
)

func Get() *http.ServeMux {
	mux := http.NewServeMux()

	uds := handler.GetUserDefineHandlers(helper.JsonParse())
	for _, ud := range uds {
		mux.Handle(ud.Path, ud.HandlerFunc)
	}

	mux.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		s := "hello"
		w.Write([]byte(s))
	})
	return mux
}
