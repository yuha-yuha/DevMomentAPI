package helper

import "encoding/json"

type JsonAPIFormat struct {
	UserDefineAPIs []UserDefineAPI `json:"apis"`
}

type UserDefineAPI struct {
	Path     string                 `json:"path"`
	Response map[string]interface{} `json:"response"`
}

func JsonParse() JsonAPIFormat {
	s :=
		`{"apis":[
		{
			"path":"/hoge",
			"response": {
				"message":"hello"
			}
		},
		{
			"path":"/def",
			"response": {
				"message":"aiueo"
			}
		}
	]}`

	sb := []byte(s)

	jaf := JsonAPIFormat{}

	json.Unmarshal(sb, &jaf)

	return jaf
}
