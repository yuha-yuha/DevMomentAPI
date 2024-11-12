package services

import (
	"encoding/json"

	"github.com/yuha-yuha/DevMomentAPI/lib"
	"github.com/yuha-yuha/DevMomentAPI/models"
)

var UserDefineAPIMap map[string][]models.UserDefineAPI

func SubscribeUserDefineAPIs(filePath string) {
	data := lib.ImportFileData(filePath)
	fileFormat := lib.FileFormat{}

	json.Unmarshal(data, &fileFormat)

	for _, api := range fileFormat.APIsFormat {
		if api.Method == "" {
			api.Method = "GET"
		}
		AddUserDefineAPI(models.UserDefineAPI{Path: api.Path, Response: api.Response, Method: api.Method, Header: api.Header})
	}

}

func FindUserDefineAPIByPath(path string) []models.UserDefineAPI {
	return UserDefineAPIMap[path]
}

func AddUserDefineAPI(ud models.UserDefineAPI) {
	if UserDefineAPIMap == nil {
		UserDefineAPIMap = make(map[string][]models.UserDefineAPI)
	}
	UserDefineAPIMap[ud.Path] = append(UserDefineAPIMap[ud.Path], ud)
}

func GetAllUserDefineAPIs() []models.UserDefineAPI {
	allApis := []models.UserDefineAPI{}
	for _, apis := range UserDefineAPIMap {
		allApis = append(allApis, apis...)
	}

	return allApis
}

func GetUserDefineAPIMap() map[string][]models.UserDefineAPI {
	return UserDefineAPIMap
}
