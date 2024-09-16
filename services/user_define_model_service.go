package services

import (
	"encoding/json"

	"github.com/yuha-yuha/DevMomentAPI/lib"
	"github.com/yuha-yuha/DevMomentAPI/models"
)

func GetUserDefineModels(filePath string) []models.UserDefineModel {
	//ファイル読み込みは関数化
	data := lib.ImportFileData(filePath)
	fileFormat := lib.FileFormat{}
	userDefineModels := []models.UserDefineModel{}

	json.Unmarshal(data, &fileFormat)

	for name, content := range fileFormat.ModelsFormat {
		userDefineModels = append(userDefineModels, models.UserDefineModel{Name: name, Content: content})
	}
	return userDefineModels
}
