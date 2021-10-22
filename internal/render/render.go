package render

import (
	"errors"
	"fmt"
	"github.com/ekateryna-tln/wallester_task/internal/config"
	"github.com/ekateryna-tln/wallester_task/internal/models"
	"github.com/justinas/nosurf"
	"html/template"
	"net/http"
	"path/filepath"
	"time"
)

var functions = template.FuncMap{
	"add":            Add,
	"formatDate":     FormatDate,
	"formatDateTime": FormatDateTime,
}

var app *config.App
var pathToTemplates = "./templates"

// SetRenderApp sets the app data for the render template package
func SetRenderApp(a *config.App) {
	app = a
}

// AddDefaultData sets default template data
func AddDefaultData(tmplData *models.TemplateData, r *http.Request) *models.TemplateData {
	tmplData.Flash = app.Session.PopString(r.Context(), "flash")
	tmplData.Error = app.Session.PopString(r.Context(), "error")
	tmplData.Warning = app.Session.PopString(r.Context(), "warning")
	tmplData.OverrideWarning = app.Session.PopString(r.Context(), "override_warning")
	fmt.Println(tmplData.OverrideWarning)
	tmplData.CSRFToken = nosurf.Token(r)
	return tmplData
}

// Add returns a and b sum result
func Add(a, b int) int {
	return a + b
}

// FormatDate returns date in yyyy-MM-dd format
func FormatDate(t time.Time) string {
	return t.Format("2006-01-02")
}

// FormatDateTime returns date in yyyy-MM-dd HH:mm:ss format
func FormatDateTime(t time.Time) string {
	return t.Format("2006-01-02 15:04:05")
}

// Template renders templates using http/template
func Template(w http.ResponseWriter, r *http.Request, tmpl string, tmplData *models.TemplateData) error {
	var templateList map[string]*template.Template
	if app.UseCache {
		templateList = app.TemplateCache
	} else {
		var err error
		templateList, err = CreateTemplateCache()
		if err != nil {
			return err
		}
	}

	template, ok := templateList[tmpl]
	if !ok {
		return errors.New("could not get template from template cache")
	}

	tmplData = AddDefaultData(tmplData, r)
	err := template.Execute(w, tmplData)
	if err != nil {
		return err
	}
	return nil
}

// CreateTemplateCache create a template cache
func CreateTemplateCache() (map[string]*template.Template, error) {
	templateCache := make(map[string]*template.Template)

	pages, err := filepath.Glob(fmt.Sprintf("%s/*.page.tmpl", pathToTemplates))
	if err != nil {
		return templateCache, err
	}

	layouts, err := filepath.Glob(fmt.Sprintf("%s/*.layout.tmpl", pathToTemplates))
	if err != nil {
		return templateCache, err
	}

	for _, page := range pages {
		name := filepath.Base(page)
		template, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			return templateCache, err
		}

		if len(layouts) > 0 {
			template, err = template.ParseGlob(fmt.Sprintf("%s/*.layout.tmpl", pathToTemplates))
			if err != nil {
				return templateCache, err
			}
		}

		templateCache[name] = template
	}
	return templateCache, nil
}
