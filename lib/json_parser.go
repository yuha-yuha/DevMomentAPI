package lib

import (
	"encoding/json"
	"io"
	"log"
	"os"
	"regexp"

	"github.com/yuha-yuha/DevMomentAPI/models"
)

type FileFormat struct {
	APIsFormat []struct {
		Path     string      `json:"path"`
		Response interface{} `json:"response"`
	} `json:"apis"`

	ModelsFormat map[string]map[string]interface{} `json:"models"`
}

type DevMomentAPIFormat struct {
	UserDefineAPIs   []UserDefineAPI                   `json:"apis"`
	UserDefineModels map[string]map[string]interface{} `json:"models"`
}

type UserDefineAPI struct {
	Path     string      `json:"path"`
	Response interface{} `json:"response"`
}

func JsonParse() DevMomentAPIFormat {

	data := ImportFileData("./sample.json")

	jaf := DevMomentAPIFormat{}

	json.Unmarshal(data, &jaf)

	return jaf
}

func GetUserDefineModels() []models.UserDefineModel {
	data := ImportFileData("./sample.json")
	fileFormat := FileFormat{}
	userDefineModels := []models.UserDefineModel{}

	json.Unmarshal(data, &fileFormat)

	for name, content := range fileFormat.ModelsFormat {
		userDefineModels = append(userDefineModels, models.UserDefineModel{Name: name, Content: content})
	}
	return userDefineModels
}

func GetUserDefineAPIs() []models.UserDefineAPI {
	data := ImportFileData("./sample.json")
	fileFormat := FileFormat{}
	userDefineAPIs := []models.UserDefineAPI{}

	json.Unmarshal(data, &fileFormat)

	for _, api := range fileFormat.APIsFormat {
		userDefineAPIs = append(userDefineAPIs, models.UserDefineAPI{Path: api.Path, Response: api.Response})
	}
	return userDefineAPIs
}
func ModelUnpackforResponseJson(dmaf *DevMomentAPIFormat) {
	for modelName, userDefineModel := range dmaf.UserDefineModels {
		for fieldName, modelField := range userDefineModel {
			ValueIsMap(&modelField, dmaf.UserDefineModels)
			userDefineModel[fieldName] = modelField
		}

		dmaf.UserDefineModels[modelName] = userDefineModel
	}
	for i := range dmaf.UserDefineAPIs {
		UserDefineAPI := &dmaf.UserDefineAPIs[i]
		ValueIsMap(&UserDefineAPI.Response, dmaf.UserDefineModels)
	}
}

func ModelUnpackforResponseJsonV2(apis []*models.UserDefineAPI, models []models.UserDefineModel) {

	for i, model := range models {
		for contentKey, contentValue := range model.Content {
			ValueIsMapV2(&contentValue, models)
			models[i].Content[contentKey] = contentValue
		}
	}

	for _, api := range apis {
		ValueIsMapV2(&(api.Response), models)
	}
}

func ValueIsMap(v *interface{}, models map[string]map[string]interface{}) {
	if m, ok := (*v).(map[string]interface{}); ok {
		for k, v := range m {
			ValueIsMap(&v, models)
			m[k] = v
		}
	} else {
		pattern := `\$\{[^}]*\}`
		re, err := regexp.Compile(pattern)
		if err != nil {
			panic(err)
		}

		if str, ok := (*v).(string); ok {
			if re.MatchString(str) {
				modelStr := str[2 : len(str)-1]
				if model, ok := models[modelStr]; ok {
					log.Println("変換しました")
					*v = model
				}
			}
		}
	}
}

func ValueIsMapV2(v *interface{}, models []models.UserDefineModel) {
	if m, ok := (*v).(map[string]interface{}); ok {
		for k, v := range m {
			ValueIsMapV2(&v, models)
			m[k] = v
		}
	} else {
		pattern := `\$\{[^}]*\}`
		re, err := regexp.Compile(pattern)
		if err != nil {
			panic(err)
		}

		if str, ok := (*v).(string); ok {
			if re.MatchString(str) {
				modelStr := str[2 : len(str)-1]
				for _, model := range models {
					if model.Name == modelStr {
						*v = model.Content
					}
				}
			}
		}
	}
}
func ImportFileData(path string) []byte {
	file, err := os.Open(path)

	if err != nil {
		log.Panic(err)
	}

	file.Stat()
	byteData, err := io.ReadAll(file)

	if err != nil {
		log.Panic(err)
	}

	return byteData
}
