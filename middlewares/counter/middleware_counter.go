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
		count := atomic.AddUint64(&counter, 1)

		log.WithField("Count", count).
			Trace("Total request(s) count")
		next.ServeHTTP(w, r)
	})
}
