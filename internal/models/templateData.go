package models

import (
	"github.com/ekateryna-tln/wallester_task/internal/forms"
)

//TemplateData holds data sent from handlers to templates
type TemplateData struct {
	Data                    map[string]interface{}
	CSRFToken               string
	Form                    *forms.Form
	Error                   string
	Flash                   string
	Warning                 string
	OverrideWarning         string
	TemplateSetup           string
	CurrentLanguage         string
	AllowedLocales          []string
	CurrentUrlWithoutLocale string
}
