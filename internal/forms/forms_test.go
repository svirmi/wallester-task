package forms

import (
	"fmt"
	"net/url"
	"reflect"
	"testing"
	"time"
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

func TestForm_IsValidDate(t *testing.T) {
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
			f.IsValidDate("birthdate")
			if got := f.Errors.Get("birthdate"); got != tt.want {
				t.Errorf("MaxLength() = %s, want %s", got, tt.want)
			}
		})
	}
}

func TestForm_IsValidAge(t *testing.T) {
	currentDate := time.Now()

	type testArgs struct {
		birthdate time.Time
		minAge    int
		maxAge    int
	}

	tests := []struct {
		name string
		args testArgs
		want string
	}{
		{"valid_age", testArgs{currentDate.AddDate(-30, 0, 0), 20, 60}, ""},
		{"invalid_age_min_limit", testArgs{currentDate.AddDate(-1, 0, 0), 20, 0}, "Age should be more than 20"},
		{"invalid_age_max_limit", testArgs{currentDate.AddDate(-200, 0, 0), 0, 100}, "Age should be less than 100"},
		{"invalid_age_limits", testArgs{currentDate.AddDate(-200, 0, 0), 20, 100}, "Age should be more than 20 and less than 100"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &Form{
				Values: url.Values{},
				Errors: errors(map[string][]string{}),
			}
			f.IsValidAge("birthdate", tt.args.birthdate, tt.args.minAge, tt.args.maxAge)
			if got := f.Errors.Get("birthdate"); got != tt.want {
				t.Errorf("MaxLength() = %s, want %s", got, tt.want)
			}
		})
	}
}

func TestForm_MaxLength(t *testing.T) {
	invalidFieldMaxLength := 3
	validFieldMaxLength := 50

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
		{"invalid_field",
			args{field: "some_field", length: invalidFieldMaxLength},
			fmt.Sprintf("This field must be less than %d characters long", invalidFieldMaxLength),
		},
		{"valid_field", args{field: "some_field", length: validFieldMaxLength}, ""},
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

func TestForm_MinLength(t *testing.T) {
	invalidFieldMinLength := 50
	validFieldMinLength := 3

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
		{"invalid_field",
			args{field: "some_field", length: invalidFieldMinLength},
			fmt.Sprintf("This field must be more than %d characters long", invalidFieldMinLength),
		},
		{"valid_field", args{field: "some_field", length: validFieldMinLength}, ""},
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
			f.MinLength(tt.args.field, tt.args.length)
			if got := f.Errors.Get(tt.args.field); got != tt.want {
				t.Errorf("MinLength() = %s, want %s", got, tt.want)
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

func TestForm_GenderIsValid(t *testing.T) {
	tests := []struct {
		name string
		args url.Values
		want string
	}{
		{"empty_gender", url.Values{}, "Please select gender"},
		{"invalid_gender", url.Values{"gender": {"invalid_gender"}}, "Please select gender"},
		{"valid_male", url.Values{"gender": {"Male"}}, ""},
		{"valid_female", url.Values{"gender": {"Female"}}, ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &Form{
				Values: tt.args,
				Errors: errors(map[string][]string{}),
			}
			f.IsValidGender("gender")
			if got := f.Errors.Get("gender"); got != tt.want {
				t.Errorf("MaxLength() = %s, want %s", got, tt.want)
			}
		})
	}
}
