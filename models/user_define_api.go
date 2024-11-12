package models

type UserDefineAPI struct {
	Path     string
	Method   string
	Header   map[string]string
	Response interface{}
}
