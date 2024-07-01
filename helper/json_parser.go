package helper

type JsonFormat struct {
	APIs []struct {
		Path     string                 `json:"path"`
		Response map[string]interface{} `json:"response"`
	} `json:"apis"`
}
