package forms

import (
	"fmt"
	"github.com/asaskevich/govalidator"
	"github.com/ekateryna-tln/wallester_task/internal/enums"
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

// MaxLength check for string minimum length
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

// IsValidBirthdate checks for valid birthdate. Checks valid age in case if min age or max age given
func (f *Form) IsValidBirthdate(field string, minAge, maxAge int) {
	layout := "2006-01-02"
	_, err := time.Parse(layout, f.Get(field))
	if err != nil {
		f.Errors.Add(field, "Invalid birthdate")
	}
}

func (f *Form) IsValidGender(field string) {
	if !enums.Exists(f.Get(field)) {
		f.Errors.Add(field, "Please select gender")
	}
}
