package paginator

import (
	"reflect"

	"github.com/astaxie/beego/orm"
)

type copyFunc func(interface{}, interface{})

// GetPage loads page of objects.
func (p *Paginator) GetPage(object interface{}, qs orm.QuerySeter, dstObj interface{}, copy copyFunc) (interface{}, int64, error) {
	// Create a slice to begin with
	var count int64
	myType := reflect.TypeOf(object)
	slice := reflect.MakeSlice(reflect.SliceOf(myType), 0, 10)
	// Create a pointer to a slice value and set it to the slice
	x := reflect.New(slice.Type())
	x.Elem().Set(slice)
	// Apply filters to QuerySet
	if p.Filters != nil {
		for filterCond, filterVal := range p.Filters {
			qs = qs.Filter(filterCond, filterVal)
		}
	}
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
		copy(s.Interface(), d.Interface())
		dstSlice = reflect.Append(dstSlice, d)
	}
	return dstSlice.Interface(), count, err
}
