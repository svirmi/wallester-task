package config

import (
	"github.com/alexedwards/scs/v2"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"html/template"
)

// App holds the application config and data
type App struct {
	UseCache                bool
	TemplateCache           map[string]*template.Template
	CookieSecure            bool
	Session                 *scs.SessionManager
	CurrentLocale           string
	Translations            interface{}
	Locales                 *i18n.Localizer
	AllowedLocales          []string
	CurrentUrlWithoutLocale string
}
