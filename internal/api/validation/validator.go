package validation

import (
	"fmt"
	"github.com/accmeboot/issueshift/internal/domain"
	"reflect"
	"strings"
)

type Validator struct {
	Errors domain.Envelope
}

func NewValidator() *Validator {
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
				err := required(val.Field(i).String(), field.Name)
				if err != nil {
					v.Errors[strings.ToLower(field.Name)] = required(val.Field(i).String(), strings.ToLower(field.Name))
				}
			}
		}

		// If filed is from embedded struct
		if field.Anonymous {
			v.Validate(val.Field(i).Interface())
		}
	}

	return len(v.Errors) == 0
}

func required(value string, name string) *string {
	if value == "" {
		message := fmt.Sprintf("field %s is required", name)
		return &message
	}

	return nil
}
