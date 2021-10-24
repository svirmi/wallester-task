package main

import (
	"fmt"
	"github.com/ekateryna-tln/wallester_task/internal/helpers"
	"github.com/justinas/nosurf"
	"net/http"
	"strings"
)

// NoSurf adds CSRF protection to all POST requests
func NoSurf(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)
	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		Secure:   app.CookieSecure,
		SameSite: http.SameSiteLaxMode,
	})
	return csrfHandler
}

// LoadSession loads and save the session on every request
func LoadSession(next http.Handler) http.Handler {
	return session.LoadAndSave(next)
}

// URLHandler adds locale if it is empty, removes slashes
func URLHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestUri := r.RequestURI
		lastChar := requestUri[len(requestUri)-1:]
		if lastChar == "/" && requestUri != "/" {
			http.Redirect(w, r, strings.TrimSuffix(r.URL.Path, "/"), http.StatusSeeOther)
			return
		}

		exploded := strings.Split(r.RequestURI, "/")
		locale := exploded[1]
		if locale != "static" {
			if !helpers.CheckValueInMap(GetAllowedLocales(), locale) || locale == "" {
				http.Redirect(w, r, fmt.Sprintf("/%s%s", defaultLocale, r.URL.Path), http.StatusSeeOther)
				return
			}
		}
		app.Locales = SetCurrentLocale(locale)
		app.CurrentLocale = GetCurrentLocale()
		app.AllowedLocales = GetAllowedLocales()
		app.CurrentUrlWithoutLocale = strings.Replace(r.URL.Path, fmt.Sprintf("/%s", locale), "", -1)
		next.ServeHTTP(w, r)
	})
}
