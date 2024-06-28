package router

import "net/http"

func Get() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		s := "hello"
		w.Write([]byte(s))
	})
	return mux
}
