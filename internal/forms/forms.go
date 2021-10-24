package forms

import (
	"fmt"
	"github.com/asaskevich/govalidator"
	"github.com/ekateryna-tln/wallester_task/internal/enums"
	strip "github.com/grokify/html-strip-tags-go"
	"net/url"
	"strings"
	"time"
)

// Form custom form struct, embeds an url.Values object
type Form struct {
	url.Values
	Errors errors
}

// New initializes a form structure
func New(data url.Values) *Form {
	return &Form{
		data,
		errors(map[string][]string{}),
	}
}

// CheckRequiredFields checks for required fields
func (f *Form) CheckRequiredFields(fields ...string) {
	for _, field := range fields {
		if strings.TrimSpace(f.Get(field)) == "" {
			f.Errors.Add(field, "This field can not be empty")
		}
	}
}

// CheckHTML checks if the code contains html
func (f *Form) CheckHTML(fields ...string) {
	for _, field := range fields {
		stripped := strip.StripTags(f.Get(field))
		if len(f.Get(field)) != len(stripped) {
			f.Errors.Add(field, "This field should not contain HTML")
		}
	}
}

// MaxLength check for string maximum length
func (f *Form) MaxLength(field string, length int) bool {
	if len(f.Get(field)) == 0 {
		f.Errors.Add(field, "This field can not be empty")
		return false
	}
	if len(f.Get(field)) > length {
		f.Errors.Add(field, fmt.Sprintf("This field must be less than %d characters long", length))
		return false
	}
	return true
}

// MinLength check for string minimum length
func (f *Form) MinLength(field string, length int) bool {
	if len(f.Get(field)) == 0 {
		f.Errors.Add(field, "This field can not be empty")
		return false
	}
	if len(f.Get(field)) < length {
		f.Errors.Add(field, fmt.Sprintf("This field must be more than %d characters long", length))
		return false
	}
	return true
}

// IsEmail checks for valid email address
func (f *Form) IsEmail(field string) {
	if !govalidator.IsEmail(f.Get(field)) {
		f.Errors.Add(field, "Invalid email address")
	}
}

// Valid returns true if there are no errors, otherwise false
func (f *Form) Valid() bool {
	return len(f.Errors) == 0
}

// IsValidDate checks for valid date
func (f *Form) IsValidDate(field string) (time.Time, bool) {
	layout := "2006-01-02"
	birthdate, err := time.Parse(layout, f.Get(field))
	if err != nil {
		f.Errors.Add(field, "Invalid birthdate")
		return time.Time{}, false
	}

	return birthdate, true
}

// IsValidAge checks valid age in case if min age or max age given
func (f *Form) IsValidAge(field string, date time.Time, minAge, maxAge int) {
	if minAge != 0 || maxAge != 0 {
		age := getAge(date)
		if minAge != 0 && maxAge != 0 && (age < minAge || age > maxAge) {
			f.Errors.Add(field, fmt.Sprintf("Age should be more than %d and less than %d", minAge, maxAge))
		} else if minAge != 0 && age < minAge {
			f.Errors.Add(field, fmt.Sprintf("Age should be more than %d", minAge))
		} else if maxAge != 0 && age > maxAge {
			f.Errors.Add(field, fmt.Sprintf("Age should be less than %d", maxAge))
		}
	}

}

// IsValidGender checks it the gender valid
func (f *Form) IsValidGender(field string) {
	if !enums.Exists(f.Get(field)) {
		f.Errors.Add(field, "Please select gender")
	}
}

// getAge returns customer age according to his birthdate
func getAge(birthday time.Time) int {
	now := time.Now()
	years := now.Year() - birthday.Year()
	if now.YearDay() < birthday.YearDay() {
		years--
	}
	return years
}
