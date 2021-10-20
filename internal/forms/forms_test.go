package forms

import (
	"fmt"
	"net/url"
	"reflect"
	"testing"
)

func TestForm_CheckRequiredFields(t *testing.T) {
	missingField := "c"
	requiredFields := []string{"a", "b", "c"}

	tests := []struct {
		name    string
		args    url.Values
		want    bool
		wantMsg string
	}{
		{"valid_field", url.Values{"a": {"a"}, "b": {"b"}, "c": {"c"}}, true, ""},
		{"invalid_field", url.Values{"a": {"a"}, "b": {"b"}}, false, "This field can not be empty"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &Form{
				Values: tt.args,
				Errors: errors(map[string][]string{}),
			}
			f.CheckRequiredFields(requiredFields[0], requiredFields[1], requiredFields[2])

			if got := f.Valid(); got != tt.want {
				t.Errorf("MaxLength() = %v, want %v", got, tt.want)
			}
			if got := f.Errors.Get(missingField); got != tt.wantMsg {
				t.Errorf("MaxLength() = %v, want %v", got, tt.wantMsg)
			}
		})
	}
}

func TestForm_IsEmail(t *testing.T) {
	tests := []struct {
		name string
		args url.Values
		want string
	}{
		{"empty_email", url.Values{}, "Invalid email address"},
		{"invalid_email", url.Values{"email": {"invalid_email"}}, "Invalid email address"},
		{"valid_email", url.Values{"email": {"valid_email@test.test"}}, ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &Form{
				Values: tt.args,
				Errors: errors(map[string][]string{}),
			}
			f.IsEmail("email")
			if got := f.Errors.Get("email"); got != tt.want {
				t.Errorf("MaxLength() = %s, want %s", got, tt.want)
			}
		})
	}
}

func TestForm_IsValidBirthdate(t *testing.T) {
	tests := []struct {
		name string
		args url.Values
		want string
	}{
		{"empty_birthdate", url.Values{}, "Invalid birthdate"},
		{"invalid_birthdate", url.Values{"birthdate": {"invalid_birthdate"}}, "Invalid birthdate"},
		{"valid_birthdate", url.Values{"birthdate": {"1995-12-24"}}, ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &Form{
				Values: tt.args,
				Errors: errors(map[string][]string{}),
			}
			f.IsValidBirthdate("birthdate", 0, 0)
			if got := f.Errors.Get("birthdate"); got != tt.want {
				t.Errorf("MaxLength() = %s, want %s", got, tt.want)
			}
		})
	}
}

func TestForm_MaxLength(t *testing.T) {
	type args struct {
		field  string
		length int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"empty_field", args{}, "This field can not be empty"},
		{"invalid_field", args{field: "some_field", length: 3}, fmt.Sprintf("This field must be less than %d characters long", 3)},
		{"valid_field", args{field: "some_field", length: 50}, ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			formData := url.Values{}
			if tt.args.field != "" {
				formData.Add(tt.args.field, tt.args.field)
			}
			f := &Form{
				Values: formData,
				Errors: errors(map[string][]string{}),
			}
			f.MaxLength(tt.args.field, tt.args.length)
			if got := f.Errors.Get(tt.args.field); got != tt.want {
				t.Errorf("MaxLength() = %s, want %s", got, tt.want)
			}
		})
	}
}

func TestForm_Valid(t *testing.T) {
	formError := errors{}
	formError.Add("first_name", "This field can not be empty")

	tests := []struct {
		name   string
		errors errors
		want   bool
	}{
		{"valid_form", map[string][]string{}, true},
		{"invalid_form", formError, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &Form{
				Values: url.Values{},
				Errors: tt.errors,
			}
			if got := f.Valid(); got != tt.want {
				t.Errorf("Valid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNew(t *testing.T) {
	form := &Form{
		url.Values{},
		map[string][]string{},
	}
	if got := New(url.Values{}); !reflect.DeepEqual(got, form) {
		t.Errorf("New() = %v, want %v", got, form)
	}
}
