package middlewares

import (
	"fmt"
	"net/http"
	"time"
)

func AccessLogger(h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		acsTime := time.Now()

		h.ServeHTTP(w, r)
		fmt.Println("endpoint path: \""+r.URL.Path+"\"  |  ", "method:", r.Method, "  |  ", "Access time:", acsTime)

	}

	return http.HandlerFunc(fn)
}
