package validationMiddleware

import (
	"crptoApi/pkg/constants"
	"encoding/json"
	"net/http"
)

type ValidationMiddleware struct {
	next http.Handler
}

func newValidationMiddleware(next http.Handler) *ValidationMiddleware {
	return &ValidationMiddleware{next: next}
}

func (v *ValidationMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	authHeader := r.Header.Get("Authorization")
	if authHeader != constants.AUTH_TOKEN {
		w.WriteHeader(http.StatusUnauthorized)
		if err := json.NewEncoder(w).Encode(constants.AUTH_ERROR); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}
	v.next.ServeHTTP(w, r)
}

func ValidationMiddlewareFunc(next http.Handler) http.Handler {
	return newValidationMiddleware(next)
}
