package paginator

import (
	"fmt"
	"reflect"

	"github.com/astaxie/beego/orm"
)

type copyFunc func(interface{}, interface{})

func makeSlice(object interface{}) reflect.Value {
	return reflect.MakeSlice(reflect.SliceOf(reflect.TypeOf(object)), 0, 10)
}

// GetPage loads page of objects.
func (p *Paginator) GetPage(object interface{}, qs orm.QuerySeter, dstObj interface{}, copy copyFunc) (intrf interface{}, count int64, err error) {
	// Create a slice to begin with
	defer func() {
		if e := recover(); e != nil {
			intrf = makeSlice(dstObj).Interface()
			err = fmt.Errorf("%v", e)
		}
	}()
	slice := makeSlice(object)
	dstType := reflect.TypeOf(dstObj)
	dstSlice := makeSlice(dstObj)
	// Create a pointer to a slice value and set it to the slice
	x := reflect.New(slice.Type())
	x.Elem().Set(slice)
	// Apply filters to QuerySet
	if p.Filters != nil {
		for filterCond, filterVal := range p.Filters {
			qs = qs.Filter(filterCond, filterVal)
		}
	}
	// is Order set?
	if p.Order != "" {
		qs = qs.OrderBy(p.Order)
	}
	// here we can get panic so we catch it in defer func
	if _, err = qs.Limit(p.Limit, p.Offset).All(x.Interface()); err == nil {
		if count, err = qs.Count(); err == nil {
			xElem := x.Elem()
			for i := 0; i < xElem.Len(); i++ {
				s := xElem.Index(i)
				d := reflect.New(dstType.Elem())
				copy(s.Interface(), d.Interface())
				dstSlice = reflect.Append(dstSlice, d)
			}
		}
	}
	return dstSlice.Interface(), count, err
}
