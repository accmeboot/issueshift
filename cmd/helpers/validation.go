package helpers

import (
	"fmt"
	"github.com/accmeboot/issueshift/internal/domain"
	"reflect"
	"regexp"
	"strings"
)

type Validator struct {
	Errors domain.Error
}

func (p *Provider) NewValidator() *Validator {
	return &Validator{Errors: make(domain.Error)}
}

func (v *Validator) Clear() {
	v.Errors = domain.Error{}
}

func (v *Validator) Validate(s any) bool {
	val := reflect.ValueOf(s)

	for i := 0; i < val.NumField(); i++ {
		field := val.Type().Field(i)
		tag := field.Tag.Get("validate")

		tags := strings.Split(tag, ",")

		for _, t := range tags {
			switch t {
			case "required":
				message := required(val.Field(i).String(), strings.ToLower(field.Name))
				if _, ok := v.Errors[strings.ToLower(field.Name)]; !ok && len(message) > 0 {
					v.Errors[strings.ToLower(field.Name)] = message
				}
			case "password":
				message := password(val.Field(i).String())
				if _, ok := v.Errors[strings.ToLower(field.Name)]; !ok && len(message) > 0 {
					v.Errors[strings.ToLower(field.Name)] = message
				}
			case "email":
				message := email(val.Field(i).String())
				if _, ok := v.Errors[strings.ToLower(field.Name)]; !ok && len(message) > 0 {
					v.Errors[strings.ToLower(field.Name)] = message
				}
			case "task_status":
				message := taskStatus(val.Field(i).String())
				if _, ok := v.Errors[strings.ToLower(field.Name)]; !ok && len(message) > 0 {
					v.Errors[strings.ToLower(field.Name)] = message
				}
			}
		}

		// If field is from embedded struct
		if field.Anonymous {
			v.Validate(val.Field(i).Interface())
		}
	}

	return len(v.Errors) == 0
}

func required(value, name string) string {
	if value == "" {
		return fmt.Sprintf("field %s is required", name)
	}

	return ""
}

func email(value string) string {
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)

	if !emailRegex.MatchString(value) {
		return "email is invalid"
	}

	return ""
}

func password(value string) string {
	// TODO: patterns like ?=. don't work with go, but you can put each check into different regex
	//passwordRegex := regexp.MustCompile(`^(?=.*[a-z])(?=.*[A-Z])(?=.*\d)(?=.*[@$!%*?&])[A-Za-z\d@$!%*?&]{8,32}$`)

	if len(value) < 8 {
		return "password is too weak"
	}

	return ""
}

func taskStatus(value string) string {
	if value != "todo" && value != "in_progress" && value != "done" {
		return "allowed values: todo, in_progress, done"
	}

	return ""
}
