package paginator

import (
	"reflect"

	"github.com/antony66/beegoutils"
	"github.com/astaxie/beego/orm"
)

// GetPage loads page of objects.
func (p *Paginator) GetPage(object interface{}, qs orm.QuerySeter, dstObj interface{}) (interface{}, int64, error) {
	// Create a slice to begin with
	var count int64
	myType := reflect.TypeOf(object)
	slice := reflect.MakeSlice(reflect.SliceOf(myType), 0, 10)
	// Create a pointer to a slice value and set it to the slice
	x := reflect.New(slice.Type())
	x.Elem().Set(slice)
	_, err := qs.OrderBy(p.Order).Limit(p.Limit, p.Offset).All(x.Interface())
	if err == nil {
		count, err = qs.Count()
	}
	dstType := reflect.TypeOf(dstObj)
	dstSlice := reflect.MakeSlice(reflect.SliceOf(dstType), 0, 10)
	xElem := x.Elem()
	for i := 0; i < xElem.Len(); i++ {
		s := xElem.Index(i)
		d := reflect.New(dstType.Elem())
		beegoutils.ReflectFields(s.Interface(), d.Interface())
		dstSlice = reflect.Append(dstSlice, d)
	}
	return dstSlice.Interface(), count, err
}
