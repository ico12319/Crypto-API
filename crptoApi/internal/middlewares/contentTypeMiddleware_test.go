package middlewares

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestContentTypeMiddleware_ServeHTTP(t *testing.T) {
	dummyHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	})

	rr := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Errorf("error in the response")
	}
	contentTypeMiddleware := ContentTypeMiddlewareFunc(dummyHandler)
	contentTypeMiddleware.ServeHTTP(rr, req)

	contentType := rr.Header().Get("Content-Type")
	expected := "application/json"
	if contentType != expected {
		t.Errorf("expected content type %s got %s", expected, contentType)
	}
}
