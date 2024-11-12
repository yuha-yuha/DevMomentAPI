package handlers

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/yuha-yuha/DevMomentAPI/models"
	"github.com/yuha-yuha/DevMomentAPI/services"
)

type UserDefineHandler struct {
	Path        string `json:"path"`
	HandlerFunc http.HandlerFunc
}

func CreateUserDefineHandler(apiMaps map[string][]models.UserDefineAPI, userDefineModels []models.UserDefineModel) []UserDefineHandler {
	userDefineHandlers := []UserDefineHandler{}
	apiPointers := []*models.UserDefineAPI{}
	for _, apis := range apiMaps {
		for _, api := range apis {
			apiPointers = append(apiPointers, &api)
		}
	}

	services.ModelUnpackforResponseJson(apiPointers, userDefineModels)

	for path, apis := range apiMaps {
		handlerFunc := func(w http.ResponseWriter, r *http.Request) {
			for _, api := range apis {
				if strings.EqualFold(api.Method, r.Method) {
					w.Header().Set("Content-Type", "application/json; charset=utf-8")

					if api.Header != nil && r.Header != nil {
						for hk, hv := range api.Header {
							if rhv, ok := r.Header[hk]; ok {
								if rhv[0] != hv {
									w.WriteHeader(400)
									return
								}
							} else {
								w.WriteHeader(400)
								return
							}
						}
					} else if api.Header != nil && r.Header == nil {
						w.WriteHeader(400)
						return
					}
					jsonenc := json.NewEncoder(w)
					jsonenc.Encode(api.Response)
				}
			}
		}

		userDefHandler := UserDefineHandler{
			Path:        path,
			HandlerFunc: handlerFunc,
		}

		userDefineHandlers = append(userDefineHandlers, userDefHandler)
	}

	return userDefineHandlers
}
