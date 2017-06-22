package beegoutils

import "reflect"

// ReflectFields copies all existing field values from source struct to destination
func ReflectFields(source interface{}, destination interface{}) {
	s := reflect.Indirect(reflect.ValueOf(source))
	d := reflect.Indirect(reflect.ValueOf(destination))
	for i := 0; i < s.NumField(); i++ {
		//log.Println(s.Type().Field(i).Name, " = ", s.Field(i))
		fieldName := s.Type().Field(i).Name
		srcValue := s.FieldByName(fieldName)
		dstValue := d.FieldByName(fieldName)
		if srcValue.IsValid() && dstValue.CanAddr() {
			if srcValue.Type() == dstValue.Type() {
				dstValue.Set(srcValue)
			}
		}
	}
}
