package helper

import (
	"encoding/json"
	"io"
	"log"
	"os"
)

type DevMomentAPIFormat struct {
	UserDefineAPIs []UserDefineAPI `json:"apis"`
}

type UserDefineAPI struct {
	Path     string                 `json:"path"`
	Response map[string]interface{} `json:"response"`
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
