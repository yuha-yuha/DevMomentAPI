package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/yuha-yuha/DevMomentAPI/lib"
	"github.com/yuha-yuha/DevMomentAPI/models"
)

type UserDefineHandler struct {
	Path        string `json:"path"`
	HandlerFunc http.HandlerFunc
}

func GetUserDefineHandlers(jaf lib.DevMomentAPIFormat) []UserDefineHandler {
	userDefHandlers := []UserDefineHandler{}
	lib.ModelUnpackforResponseJson(&jaf)

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

func CreateUserDefineHandler(apis []models.UserDefineAPI, userDefineModels []models.UserDefineModel) []UserDefineHandler {
	userDefineHandlers := []UserDefineHandler{}
	apiPointers := []*models.UserDefineAPI{}
	for _, api := range apis {
		apiPointers = append(apiPointers, &api)
	}
	lib.ModelUnpackforResponseJsonV2(apiPointers, userDefineModels)

	for _, api := range apiPointers {
		handlerFunc := func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json; charset=utf-8")
			jsonenc := json.NewEncoder(w)
			jsonenc.Encode(api.Response)
		}

		userDefHandler := UserDefineHandler{
			Path:        api.Path,
			HandlerFunc: handlerFunc,
		}

		userDefineHandlers = append(userDefineHandlers, userDefHandler)
	}

	return userDefineHandlers
}
