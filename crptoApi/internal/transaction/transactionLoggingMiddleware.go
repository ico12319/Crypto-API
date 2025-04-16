package transaction

import (
	"crptoApi/internal/responseWriter"
	"log"
	"net/http"
	"time"
)

type LoggingMiddleware struct {
	next http.Handler
}

func newLoggingMiddleware(next http.Handler) *LoggingMiddleware {
	return &LoggingMiddleware{next: next}
}

func (l *LoggingMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	log.Printf("[Request] %s %s from %s",
		r.Method,
		r.RequestURI,
		r.RemoteAddr,
	)
	decoratorResponseWriter := responseWriter.NewLoggingResponseWriter(w)
	l.next.ServeHTTP(decoratorResponseWriter, r)
	duration := time.Since(start)
	log.Printf("[Response] Completed in %v with status code %d", duration, decoratorResponseWriter.GetStatusCode())
	if decoratorResponseWriter.GetStatusCode() >= http.StatusBadRequest {
		log.Printf("[Eroor] Reqeust %s %s returned error status %d", r.Method, r.RequestURI, decoratorResponseWriter.GetStatusCode())
	}
}

func LoggingMiddlewareFunc(next http.Handler) http.Handler {
	return newLoggingMiddleware(next)
}
