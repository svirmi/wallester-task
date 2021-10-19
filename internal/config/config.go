package config

import (
	"github.com/alexedwards/scs/v2"
	"html/template"
)

// App holds the application config and data
type App struct {
	UseCache      bool
	TemplateCache map[string]*template.Template
	CookieSecure  bool
	Session       *scs.SessionManager
}
