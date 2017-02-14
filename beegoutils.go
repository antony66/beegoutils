package beegoutils

import (
	"log"
	"reflect"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/validation"
	raven "github.com/getsentry/raven-go"
)

// JSONResultContainer is a container for JSON answer to GET/POST/PUT/DELETE successful calls
type JSONResultContainer struct {
	Result interface{}
}

// Paginator struct
type Paginator struct {
	Entities []interface{}
	Count    int64
	Offset   int
	Limit    int
}

// ExtendedController adds extra methods to beego.Controller
type ExtendedController struct {
	beego.Controller
}

// FinishTransaction commits or rollbacks current transaction depending on error state
func FinishTransaction(o orm.Ormer, err error) {
	if err == nil {
		o.Commit()
		return
	}
	o.Rollback()
}

// ReflectFields copies all existing field values from source struct to destination
func ReflectFields(source interface{}, destination interface{}) {
	s := reflect.ValueOf(source).Elem()
	d := reflect.ValueOf(destination).Elem()
	for i := 0; i < s.NumField(); i++ {
		//log.Println(s.Type().Field(i).Name, " = ", s.Field(i))
		d.FieldByName(s.Type().Field(i).Name).Set(s.Field(i))
	}
}

// JSONErrorWithCode returns json-encoded error with http code
func (c *ExtendedController) JSONErrorWithCode(err error, code int) {
	log.Println(err.Error())
	if code == 500 && beego.BConfig.RunMode != "dev" {
		raven.SetHttpContext(raven.NewHttp(c.Ctx.Request))
		raven.CaptureErrorAndWait(err, nil)
	}
	m := make(map[string]interface{})
	m["Errors"] = map[string]string{"Message": err.Error()}
	c.Data["json"] = m
	c.Ctx.ResponseWriter.WriteHeader(code)
	c.ServeJSON()
	c.StopRun()
}

// JSONErrorsMapWithCode returns json-encoded slice of errors with http code
func (c *ExtendedController) JSONErrorsMapWithCode(err map[string]*validation.Error, code int) {
	m := make(map[string]interface{})
	m["Errors"] = err
	c.Data["json"] = m
	c.Ctx.ResponseWriter.WriteHeader(code)
	c.ServeJSON()
	c.StopRun()
}
