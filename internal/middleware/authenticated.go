package middleware

import (
	"net/http"

	"github.com/sankalpmukim/url-shortener-go/pkg/logs"
)

func Authenticated(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, err := r.Cookie("auth")
		if err != nil {
			// log the error
			logs.Warn("Redirecting because:", err)

			http.Redirect(w, r, "/auth/login", http.StatusSeeOther)
			return
		}
		next.ServeHTTP(w, r)
	})
}
