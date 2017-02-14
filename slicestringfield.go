package beegoutils

import (
	"fmt"
	"strings"

	"github.com/astaxie/beego/orm"
)

// SliceStringField represents []string for beego ORM
type SliceStringField []string

// Value ...
func (e SliceStringField) Value() []string {
	return []string(e)
}

// Set ...
func (e *SliceStringField) Set(d []string) {
	*e = SliceStringField(d)
}

// Add ...
func (e *SliceStringField) Add(v string) {
	*e = append(*e, v)
}

// String ...
func (e *SliceStringField) String() string {
	return strings.Join(e.Value(), "\n")
}

// FieldType ...
func (e *SliceStringField) FieldType() int {
	return orm.TypeCharField
}

// SetRaw ...
func (e *SliceStringField) SetRaw(value interface{}) error {
	switch d := value.(type) {
	case []string:
		e.Set(d)
	case string:
		if len(d) > 0 {
			parts := strings.Split(d, "\n")
			v := make([]string, 0, len(parts))
			for _, p := range parts {
				v = append(v, strings.TrimSpace(p))
			}
			e.Set(v)
		}
	default:
		return fmt.Errorf("<SliceStringField.SetRaw> unknown value `%v`", value)
	}
	return nil
}

// RawValue ...
func (e *SliceStringField) RawValue() interface{} {
	return e.String()
}

// Clean ...
func (e *SliceStringField) Clean() error {
	return nil
}

var _ orm.Fielder = new(SliceStringField)
