package validation

import (
	"encoding/json"
	"errors"
	"reflect"
	"strings"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/johngtrs/go-rest-api/httperror"
)

type ApiError struct {
	Param   string `json:"param"`
	Message string `json:"message"`
}

// Return a custom error message for a specific tag
func msgForTag(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return "This field is required"
	case "email":
		return "Invalid email"
	}
	return fe.Error()
}

// Init Go Playground validator settings
func InitValidator() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterTagNameFunc(func(fld reflect.StructField) string {
			name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
			if name == "-" {
				return ""
			}
			return name
		})
	}
}

// Return the JSON response object error
func ErrorMessages(err error) map[string]interface{} {
	var ve validator.ValidationErrors
	// Check if the errors are type of ValidationErrors
	if errors.As(err, &ve) {
		out := make([]ApiError, len(ve))
		for i, fe := range ve {
			out[i] = ApiError{fe.Field(), msgForTag(fe)}
		}
		return map[string]interface{}{"error": out}
	}

	// Handle UnmarshalTypeError
	var ute *json.UnmarshalTypeError
	if errors.As(err, &ute) {
		return map[string]interface{}{"error": ute.Error()}
	}

	// Default error
	return map[string]interface{}{"error": httperror.ErrBadRequest.Error()}
}
