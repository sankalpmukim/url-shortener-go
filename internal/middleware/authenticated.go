package middleware

import (
	"net/http"

	"github.com/sankalpmukim/url-shortener-go/internal/controllers"
	"github.com/sankalpmukim/url-shortener-go/internal/cookies"
	"github.com/sankalpmukim/url-shortener-go/pkg/logs"
)

func Authenticated(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		v, err := r.Cookie("auth")
		if err != nil {
			// log the error
			logs.Warn("Redirecting because:", err)

			// tell user error
			cookies.CreateOrAppendFlash(w, r, cookies.ERROR, "Auth cookie not found")

			http.Redirect(w, r, "/auth/login", http.StatusSeeOther)
			return
		}

		// verify the cookie and get user's email
		email, err := controllers.GetEmailFromPayload(v.Value)

		if err != nil {
			// log the error
			logs.Warn("Redirecting because:", err)

			// delete the user's auth cookie
			dc := cookies.DeleteCookieCookie("auth")
			http.SetCookie(w, dc)

			// tell user error
			cookies.CreateOrAppendFlash(w, r, cookies.ERROR, "Auth cookie invalid")

			http.Redirect(w, r, "/auth/login", http.StatusSeeOther)
		}

		// set the user's email in the request context
		ctx := controllers.SetEmailInContext(r, email)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
