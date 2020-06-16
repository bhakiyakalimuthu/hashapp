package middleware

import (
	"encoding/json"
	"net/http"
)

// IsUserAllowed will verify the api key token validation
func IsUserAllowed(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if isAPIKeyValid(r) {
			next.ServeHTTP(w, r)
			return
		}
		unAuthorized(w)
	})
}

func isAPIKeyValid(r *http.Request) bool {
	if apiKey := r.Header.Get("Authorization"); len(apiKey) >= 8 && apiKey[:7] == "Bearer " {
		return APIKey == apiKey[7:]
	}
	return false
}

func unAuthorized(w http.ResponseWriter) {
	code := http.StatusUnauthorized
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(code)
	if err := json.NewEncoder(w).Encode(http.StatusText(code)); err != nil {
		http.Error(w, "", http.StatusInternalServerError)
	}
}



