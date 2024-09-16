package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/yuha-yuha/DevMomentAPI/models"
	"github.com/yuha-yuha/DevMomentAPI/services"
)

type UserDefineHandler struct {
	Path        string `json:"path"`
	HandlerFunc http.HandlerFunc
}

func CreateUserDefineHandler(apis []models.UserDefineAPI, userDefineModels []models.UserDefineModel) []UserDefineHandler {
	userDefineHandlers := []UserDefineHandler{}
	apiPointers := []*models.UserDefineAPI{}
	for _, api := range apis {
		apiPointers = append(apiPointers, &api)
	}
	services.ModelUnpackforResponseJson(apiPointers, userDefineModels)

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
