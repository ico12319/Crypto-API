package middlewares

import (
	"crptoApi/pkg/constants"
	"net/http"
)

type ContentTypeMiddleware struct {
	next http.Handler
}

func newContentTypeMiddleware(next http.Handler) *ContentTypeMiddleware {
	return &ContentTypeMiddleware{next: next}
}

func (c *ContentTypeMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(constants.CONTENT_TYPE, constants.JSON)
	c.next.ServeHTTP(w, r)
}
func ContentTypeMiddlewareFunc(next http.Handler) http.Handler {
	return newContentTypeMiddleware(next)
}
