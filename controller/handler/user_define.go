package handler

import (
	"encoding/json"
	"net/http"

	"github.com/yuha-yuha/DevMomentAPI/helper"
)

type UserDefineHandler struct {
	Path        string `json:"path"`
	HandlerFunc http.HandlerFunc
}

func GetUserDefineHandlers(jaf helper.JsonAPIFormat) []UserDefineHandler {
	userDefHandlers := []UserDefineHandler{}
	for _, UserDefineAPI := range jaf.UserDefineAPIs {
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
