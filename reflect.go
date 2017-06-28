package beegoutils

import "reflect"

// ReflectFields copies all existing field values from source struct to destination
func ReflectFields(source interface{}, destination interface{}) {
	s := reflect.Indirect(reflect.ValueOf(source))
	d := reflect.Indirect(reflect.ValueOf(destination))
	for i := 0; i < s.NumField(); i++ {
		fieldName := s.Type().Field(i).Name
		srcValue := s.FieldByName(fieldName)
		dstValue := d.FieldByName(fieldName)
		if srcValue.IsValid() && dstValue.CanAddr() {
			switch dstValue.Kind() {
			case reflect.Struct, reflect.Ptr:
				continue // just skip these fields
			default:
				dstValue.Set(srcValue)
			}
		}
	}
}
