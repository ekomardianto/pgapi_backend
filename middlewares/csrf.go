package middlewares

import (
	"crypto/rand"
	"encoding/hex"
	"net/http"
)

// GenerateCSRFToken membuat CSRF token acak
func GenerateCSRFToken() (string, error) {
	token := make([]byte, 16)
	if _, err := rand.Read(token); err != nil {
		return "", err
	}
	return hex.EncodeToString(token), nil
}

// CSRFProtection adalah middleware yang memvalidasi CSRF token
func CSRFProtection(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Hanya periksa pada metode POST, PUT, DELETE
		if r.Method == "POST" || r.Method == "PUT" || r.Method == "DELETE" {
			csrfToken := r.Header.Get("X-CSRF-Token")
			// Dapatkan token dari cookie (disetel saat generate)
			cookie, err := r.Cookie("csrf_token")
			if err != nil || csrfToken != cookie.Value {
				http.Error(w, "Invalid CSRF token", http.StatusForbidden)
				return
			}
		}
		next.ServeHTTP(w, r)
	})
}
