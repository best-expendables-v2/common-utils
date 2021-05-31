package validation

import (
	"reflect"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	validator "gopkg.in/go-playground/validator.v9"
)

func Test_ParseValidationErr(t *testing.T) {
	validate := initValidate()
	testObject := validateTest{}
	err := validate.Struct(testObject)
	vErrs := ParseValidationErr(err)
	assert.Equal(t, 3, len(vErrs))
	vErr := vErrs[0]
	assert.Equal(t, "name", vErr.Field)
	assert.Equal(t, "required", vErr.Code)
	assert.Equal(t, "", vErr.Argument)
	assert.Equal(t, "The name field is required.", vErr.Message)

	vErr = vErrs[1]
	assert.Equal(t, "score", vErr.Field)
	assert.Equal(t, "min", vErr.Code)
	assert.Equal(t, "5", vErr.Argument)
	assert.Equal(t, "The score field must be at least 5.", vErr.Message)

	vErr = vErrs[2]
	assert.Equal(t, "type", vErr.Field)
	assert.Equal(t, "customType", vErr.Code)
	assert.Equal(t, "", vErr.Argument)
	assert.Equal(t, "Validation failed for type", vErr.Message)
}

func initValidate() *validator.Validate {
	validate := validator.New()
	validate.SetTagName("validate")
	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]

		if name == "-" {
			return ""
		}

		return name
	})
	validate.RegisterValidation("customType", validateType)
	return validate
}

type validateTest struct {
	Name  string `json:"name" validate:"required,min=5"`
	Score int    `json:"score" validate:"min=5"`
	Type  string `json:"type" validate:"customType"`
}

func validateType(fl validator.FieldLevel) bool {
	if fl.Field().String() == "test" {
		return true
	}
	return false
}
