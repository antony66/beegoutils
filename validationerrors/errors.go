package validationerrors

import (
	"github.com/antony66/beegoutils"
	"github.com/astaxie/beego/validation"
)

// LoadErrors fills ValidationErrors with validation.Validation errors
func (e *ValidationErrors) LoadErrors(valid *validation.Validation) {
	e.Errors = make(map[string]*ValidationErrors_ValidationError)
	for k, v := range valid.ErrorMap() {
		ee := new(ValidationErrors_ValidationError)
		beegoutils.ReflectFields(v, ee)
		e.Errors[k] = ee
	}
}
