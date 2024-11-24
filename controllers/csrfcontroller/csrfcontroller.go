package csrfcontroller

import (
	"eko/api-pg-bpr/middlewares"
	"net/http"
)

func GenerateCSRFToken(w http.ResponseWriter, r *http.Request) {
	// Buat CSRF token
	token, err := middlewares.GenerateCSRFToken()
	if err != nil {
		http.Error(w, "Failed to generate CSRF token", http.StatusInternalServerError)
		return
	}

	// Simpan token dalam cookie
	http.SetCookie(w, &http.Cookie{
		Name:     "csrf_token",
		Value:    token,
		HttpOnly: true,
		Secure:   true, // Gunakan Secure jika di HTTPS
		SameSite: http.SameSiteStrictMode,
	})

	// Kirim token sebagai respons (opsional, hanya untuk konfirmasi)
	w.Write([]byte(token))
}
