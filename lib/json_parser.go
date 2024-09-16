package lib

import (
	"encoding/json"
	"io"
	"log"
	"os"
)

type FileFormat struct {
	APIsFormat []struct {
		Path     string      `json:"path"`
		Response interface{} `json:"response"`
	} `json:"apis"`

	ModelsFormat map[string]map[string]interface{} `json:"models"`
}

type DevMomentAPIFormat struct {
	UserDefineAPIs   []UserDefineAPIFormat             `json:"apis"`
	UserDefineModels map[string]map[string]interface{} `json:"models"`
}

type UserDefineAPIFormat struct {
	Path     string      `json:"path"`
	Response interface{} `json:"response"`
}

func JsonParse() DevMomentAPIFormat {

	data := ImportFileData("./sample.json")

	jaf := DevMomentAPIFormat{}

	json.Unmarshal(data, &jaf)

	return jaf
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
