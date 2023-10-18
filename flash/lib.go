package flash

import (
	"encoding/base64"
	"net/http"
	"strings"
	"time"
)

func SetFlash(w http.ResponseWriter, name string, value []byte) {
	c := &http.Cookie{
		Name:     name,
		Value:    encode(value),
		Path:     "/",                      // makes sure it's available for the whole domain
		Domain:   "",                       // leave it empty to default to the domain of the calling script
		MaxAge:   3600,                     // 1 hour in seconds, 0 means no 'Max-Age' attribute set. If negative, delete cookie now.
		Secure:   false,                    // true if you only want to send the cookie over HTTPS
		HttpOnly: true,                     // true if you want to prevent JavaScript access to the cookie
		SameSite: http.SameSiteDefaultMode, // or http.SameSiteLaxMode, http.SameSiteStrictMode, http.SameSiteNoneMode
	}
	http.SetCookie(w, c)
}

func GetFlash(w http.ResponseWriter, r *http.Request, name string) ([]byte, error) {
	c, err := r.Cookie(name)
	if err != nil {
		switch err {
		case http.ErrNoCookie:
			return nil, nil
		default:
			return nil, err
		}
	}
	value, err := decode(c.Value)
	if err != nil {
		return nil, err
	}

	// clearning the cookie
	dc := &http.Cookie{
		Name:     name,
		MaxAge:   -1,
		Expires:  time.Unix(1, 0),
		Path:     "/",                      // makes sure it's available for the whole domain
		Domain:   "",                       // leave it empty to default to the domain of the calling script
		Secure:   false,                    // true if you only want to send the cookie over HTTPS
		HttpOnly: true,                     // true if you want to prevent JavaScript access to the cookie
		SameSite: http.SameSiteDefaultMode, // or http.SameSiteLaxMode, http.SameSiteStrictMode, http.SameSiteNoneMode
	}
	http.SetCookie(w, dc)
	return value, nil
}

func SplitByCommas(s string) []string {
	if s == "" {
		return nil
	}
	return strings.Split(s, ",")
}

func encode(src []byte) string {
	return base64.URLEncoding.EncodeToString(src)
}

func decode(src string) ([]byte, error) {
	return base64.URLEncoding.DecodeString(src)
}
