package beegoutils

import (
	"fmt"
	"reflect"
)

// ReflectFields copies all existing field values from source struct to destination
func ReflectFields(source interface{}, destination interface{}) {
	s := reflect.Indirect(reflect.ValueOf(source))
	d := reflect.Indirect(reflect.ValueOf(destination))
	for i := 0; i < s.NumField(); i++ {
		fieldName := s.Type().Field(i).Name
		srcValue := s.FieldByName(fieldName)
		dstValue := d.FieldByName(fieldName)
		if srcValue.IsValid() && dstValue.CanAddr() {
			if srcValue.Type() == dstValue.Type() {
				dstValue.Set(srcValue)
			} else {
				switch dstValue.Kind() {
				case reflect.Struct, reflect.Ptr:
					continue // just skip these fields
				case reflect.String:
					dstValue.Set(reflect.ValueOf(fmt.Sprintf("%v", srcValue.Interface())))
				default:
					dstValue.Set(srcValue)
				}
			}
		}
	}
}
