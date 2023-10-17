package middleware

import (
	"net/http"
)

func Authenticated(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, err := r.Cookie("auth")
		if err != nil {
			http.Redirect(w, r, "/auth/login", http.StatusSeeOther)
			return
		}
		next.ServeHTTP(w, r)
	})
}
