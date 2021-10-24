package main

import (
	"encoding/json"
	"fmt"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

const defaultLocale = "en"

var bundle *i18n.Bundle
var currentLocale = ""
var pathToLocales = "./static/locales"
var allowedLocales = []string{"en", "de"}

// GetAllowedLocales returns a bunch of allowed locales
func GetAllowedLocales() []string {
	return allowedLocales
}

// GetCurrentLocale returns a current locale
func GetCurrentLocale() string {
	return currentLocale
}

// InitLocaleBundle initialises app locales bundle
func InitLocaleBundle() {
	bundle = i18n.NewBundle(language.English)
	bundle.RegisterUnmarshalFunc("json", json.Unmarshal)
	_, err := bundle.LoadMessageFile(fmt.Sprintf("%s/en.json", pathToLocales))
	if err != nil {
		fmt.Println(err)
	}
	_, err = bundle.LoadMessageFile(fmt.Sprintf("%s/de.json", pathToLocales))
	if err != nil {
		fmt.Println(err)
	}
}

// SetCurrentLocale sets current locale
func SetCurrentLocale(locale string) *i18n.Localizer {
	currentLocale = locale
	if currentLocale == "" {
		currentLocale = defaultLocale
	}

	allowedLanguages := map[string]string{
		"en": language.English.String(),
		"de": language.German.String(),
	}
	currentLanguage, ok := allowedLanguages[currentLocale]
	if !ok {
		currentLanguage = language.English.String()
	}

	return i18n.NewLocalizer(bundle, currentLanguage)

}
