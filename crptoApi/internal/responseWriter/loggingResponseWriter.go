package responseWriter

import "net/http"

type LoggingResponseWriter struct {
	http.ResponseWriter
	StatusCode int
}

func NewLoggingResponseWriter(w http.ResponseWriter) *LoggingResponseWriter {
	return &LoggingResponseWriter{ResponseWriter: w, StatusCode: http.StatusOK}
}

func (l *LoggingResponseWriter) WriteHeader(statusCode int) {
	l.StatusCode = statusCode
	l.ResponseWriter.WriteHeader(statusCode)
}
