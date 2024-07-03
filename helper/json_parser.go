package helper

import (
	"encoding/json"
	"io"
	"log"
	"os"
	"regexp"
)

type DevMomentAPIFormat struct {
	UserDefineAPIs   []UserDefineAPI                   `json:"apis"`
	UserDefineModels map[string]map[string]interface{} `json:"models"`
}

/*
 */
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

func ModelIntoResponseJson(dmaf *DevMomentAPIFormat) {
	for i := range dmaf.UserDefineAPIs {
		UserDefineAPI := &dmaf.UserDefineAPIs[i]
		ValueIsMap(&UserDefineAPI.Response, dmaf.UserDefineModels)
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
