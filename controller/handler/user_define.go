package handler

import "net/http"

type UserDefine struct {
	Path    string
	Handler http.Handler
}
