package cookies

import (
	"encoding/base64"
	"net/http"
	"strings"
	"time"

	"github.com/sankalpmukim/url-shortener-go/pkg/logs"
)

func SetFlash(w http.ResponseWriter, name string, value []byte) {
	c := CreateCookie(name, encode(value))
	http.SetCookie(w, &c)
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

func DeleteCookieCookie(name string) *http.Cookie {
	return &http.Cookie{
		Name:     name,
		MaxAge:   -1,
		Expires:  time.Unix(1, 0),
		Path:     "/",                      // makes sure it's available for the whole domain
		Domain:   "",                       // leave it empty to default to the domain of the calling script
		Secure:   false,                    // true if you only want to send the cookie over HTTPS
		HttpOnly: true,                     // true if you want to prevent JavaScript access to the cookie
		SameSite: http.SameSiteDefaultMode, // or http.SameSiteLaxMode, http.SameSiteStrictMode, http.SameSiteNoneMode
	}
}

func CreateCookie(name string, value string) http.Cookie {
	return http.Cookie{
		Name:     name,
		Value:    value,
		Path:     "/",                      // makes sure it's available for the whole domain
		Domain:   "",                       // leave it empty to default to the domain of the calling script
		MaxAge:   3600,                     // 1 hour in seconds, 0 means no 'Max-Age' attribute set. If negative, delete cookie now.
		Secure:   false,                    // true if you only want to send the cookie over HTTPS
		HttpOnly: true,                     // true if you want to prevent JavaScript access to the cookie
		SameSite: http.SameSiteDefaultMode, // or http.SameSiteLaxMode, http.SameSiteStrictMode, http.SameSiteNoneMode
	}
}

func SplitByCommas(s string) []string {
	if s == "" {
		return nil
	}
	return strings.Split(s, ",")
}

// split multiple statements by ";"
func CreateOrAppendFlash(w http.ResponseWriter, r *http.Request, name, value string) {
	// get cookie using GetFlash
	oldFlash, err := GetFlash(w, r, name)
	if err != nil {
		logs.Error("Failed to parse form(flash cookie)", err)
		http.Error(w, "Failed to parse form(flash cookie)", http.StatusInternalServerError)
		return
	}

	flash := string(oldFlash)

	// create the string
	if flash != "" {
		flash += ";"
	}
	flash += value

	// set the cookie
	SetFlash(w, name, []byte(flash))
}

func encode(src []byte) string {
	return base64.URLEncoding.EncodeToString(src)
}

func decode(src string) ([]byte, error) {
	return base64.URLEncoding.DecodeString(src)
}

type FlashInfo struct {
	Success []string
	Info    []string
	Error   []string
}

func GetFlashInfo(w http.ResponseWriter, r *http.Request) (FlashInfo, error) {
	successBytes, err := GetFlash(w, r, SUCCESS)
	if err != nil {
		http.Error(w, "Failed to parse form(flash cookie)", http.StatusInternalServerError)
		return FlashInfo{}, err
	}
	errorBytes, err := GetFlash(w, r, ERROR)
	if err != nil {
		http.Error(w, "Failed to parse form(flash cookie)", http.StatusInternalServerError)
		return FlashInfo{}, err
	}
	infoBytes, err := GetFlash(w, r, INFO)
	if err != nil {
		http.Error(w, "Failed to parse form(flash cookie)", http.StatusInternalServerError)
		return FlashInfo{}, err
	}
	success := strings.Split(string(successBytes), ";")
	errors := strings.Split(string(errorBytes), ";")
	info := strings.Split(string(infoBytes), ";")
	return FlashInfo{
		Success: cleanString(success),
		Error:   cleanString(errors),
		Info:    cleanString(info),
	}, nil
}

// trim spaces and newlines
// remove empty strings
func cleanString(val []string) []string {
	var res []string
	for _, v := range val {
		v = strings.TrimSpace(v)
		if v != "" {
			res = append(res, v)
		}
	}
	return res
}

var SUCCESS = "success"
var INFO = "info"
var ERROR = "error"
