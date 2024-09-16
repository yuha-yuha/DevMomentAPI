package services

import (
	"encoding/json"

	"github.com/yuha-yuha/DevMomentAPI/lib"
	"github.com/yuha-yuha/DevMomentAPI/models"
)

func GetUserDefineAPIs(filePath string) []models.UserDefineAPI {
	data := lib.ImportFileData(filePath)
	fileFormat := lib.FileFormat{}
	userDefineAPIs := []models.UserDefineAPI{}

	json.Unmarshal(data, &fileFormat)

	for _, api := range fileFormat.APIsFormat {
		userDefineAPIs = append(userDefineAPIs, models.UserDefineAPI{Path: api.Path, Response: api.Response})
	}
	return userDefineAPIs
}
