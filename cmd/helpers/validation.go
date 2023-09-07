package helpers

import (
	"fmt"
	"github.com/accmeboot/issueshift/internal/domain"
	"reflect"
	"regexp"
	"strings"
)

type Validator struct {
	Errors domain.Envelope
}

func (p *Provider) NewValidator() *Validator {
	return &Validator{Errors: make(domain.Envelope)}
}

func (v *Validator) Clear() {
	v.Errors = domain.Envelope{}
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
				err := required(val.Field(i).String(), strings.ToLower(field.Name))
				if _, ok := v.Errors[strings.ToLower(field.Name)]; !ok && err != nil {
					v.Errors[strings.ToLower(field.Name)] = err
				}
			case "password":
				err := password(val.Field(i).String())
				if _, ok := v.Errors[strings.ToLower(field.Name)]; !ok && err != nil {
					v.Errors[strings.ToLower(field.Name)] = err
				}
			case "email":
				err := email(val.Field(i).String())
				if _, ok := v.Errors[strings.ToLower(field.Name)]; !ok && err != nil {
					v.Errors[strings.ToLower(field.Name)] = err
				}
			case "task_status":
				err := taskStatus(val.Field(i).String())
				if _, ok := v.Errors[strings.ToLower(field.Name)]; !ok && err != nil {
					v.Errors[strings.ToLower(field.Name)] = err
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

func required(value, name string) *string {
	if value == "" {
		message := fmt.Sprintf("field %s is required", name)
		return &message
	}

	return nil
}

func email(value string) *string {
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)

	if !emailRegex.MatchString(value) {
		message := fmt.Sprint("email is invalid")
		return &message
	}

	return nil
}

func password(value string) *string {
	// TODO: patterns like ?=. don't work with go, but you can put each check into different regex
	//passwordRegex := regexp.MustCompile(`^(?=.*[a-z])(?=.*[A-Z])(?=.*\d)(?=.*[@$!%*?&])[A-Za-z\d@$!%*?&]{8,32}$`)

	if len(value) < 8 {
		message := fmt.Sprint("password is too weak")
		return &message
	}

	return nil
}

func taskStatus(value string) *string {
	if value != "todo" && value != "in_progress" && value != "done" {
		message := "allowed values: todo, in_progress, done"
		return &message
	}

	return nil
}
