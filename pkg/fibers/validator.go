package fibers

import (
	"errors"
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
)

var v = validator.New()

func Validate(data any) error {
	results := []string{}
	if errs := v.Struct(data); errs != nil {
		for _, err := range errs.(validator.ValidationErrors) {
			if err.Param() == "" {
				results = append(results, fmt.Sprintf("[%s:%s]: '%s'", err.Field(), err.Value(), err.Tag()))
				continue
			}
			results = append(results, fmt.Sprintf("[%s:%s]: '%s-%s'", err.Field(), err.Value(), err.Tag(), err.Param()))
		}
	}
	if len(results) > 0 {
		return errors.New(strings.Join(results, "\n"))
	}
	return nil
}
