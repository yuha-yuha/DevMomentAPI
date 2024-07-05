package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/yuha-yuha/DevMomentAPI/helper"
)

type UserDefineHandler struct {
	Path        string `json:"path"`
	HandlerFunc http.HandlerFunc
}

func GetUserDefineHandlers(jaf helper.DevMomentAPIFormat) []UserDefineHandler {
	userDefHandlers := []UserDefineHandler{}
	helper.ModelUnpackforResponseJson(&jaf)

	for _, UserDefineAPI := range jaf.UserDefineAPIs {
		log.Println(UserDefineAPI.Response)
		handlerFunc := func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json; charset=utf-8")
			jsonenc := json.NewEncoder(w)
			jsonenc.Encode(UserDefineAPI.Response)
		}

		userDefHandler := UserDefineHandler{
			Path:        UserDefineAPI.Path,
			HandlerFunc: handlerFunc,
		}

		userDefHandlers = append(userDefHandlers, userDefHandler)
	}

	return userDefHandlers
}
