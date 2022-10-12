package validators

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

type ValidationError struct {
	Field  string `json:"field"`
	Reason string `json:"reason"`
}

func Bind(c *gin.Context, obj interface{}) error {
	b := binding.Default(c.Request.Method, c.ContentType())
	ret := c.ShouldBindWith(obj, b)
	return ret
}

// Warp the error info in a object
func NewError(params string, err error) []ValidationError {
	res := []ValidationError{}
	res = append(res, ValidationError{Field: params, Reason: err.Error()})
	return res
}
func NewValidatorError(err validator.ValidationErrors) []ValidationError {
	errs := []ValidationError{}

	for _, f := range err {
		errs = append(errs, ValidationError{Field: f.Field(), Reason: ValidationErrorToText(f)})
	}
	return errs
}
func ValidationErrorToText(errValidation validator.FieldError) string {
	name := errValidation.Field()

	switch errValidation.Tag() {
	case "required":
		return fmt.Sprintf("%s is required", name)
	case "max":
		return fmt.Sprintf("%s can't be longer than %s", name, errValidation.Param())
	case "min":
		return fmt.Sprintf("%s must be longer than %s", name, errValidation.Param())
	case "email":
		return "Invalid email format"
	case "emailExist":
		return "Email already taken"
	case "usernameExist":
		return "Username already taken"
	case "eqfield":
		return fmt.Sprintf("%s must same with %s", name, errValidation.Param())
	case "gte":
		return fmt.Sprintf("%s must Greather than or equal %s", name, errValidation.Param())
	default:
		return fmt.Sprintf("%s is not valid input", name)
	}
}
