package middleware

import (
	"net/http"
	"sync/atomic"

	log "github.com/sirupsen/logrus"
)

var (
	counter uint64
)

func MiddlewareCounter(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		newCounter := atomic.AddUint64(&counter, 1)

		log.Info(newCounter, " request(s) since the start of the app.")
		next.ServeHTTP(w, r)
	})
}
