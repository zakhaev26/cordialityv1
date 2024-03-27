package middlewares

import (
	"net/http"
	"time"

	"github.com/charmbracelet/log"
)

func LoggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		log.Info("LOGS", "METHOD",r.Method, "URI", r.RequestURI, "LATENCY", time.Since(start))
	})
}
