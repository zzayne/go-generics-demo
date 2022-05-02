package reflect

import (
	"reflect"
)

func GetStructFieldNames(v interface{}) []string {
	var res []string
	rt := reflect.TypeOf(v)
	n := rt.NumField()
	for i := 0; i < n; i++ {
		res = append(res, rt.Field(i).Name)
	}
	return res
}

func GetStructFieldTypes(v interface{}) []string {
	var res []string
	rt := reflect.TypeOf(v)
	n := rt.NumField()
	for i := 0; i < n; i++ {
		res = append(res, rt.Field(i).Type.Name())
	}
	return res
}

func SetFieldValue(v interface{}, filed string, val interface{}) error {
	rv := reflect.ValueOf(v)
	if rv.Kind() != reflect.Ptr {
		return ErrShouldBePrt
	}
	elem := rv.Elem()
	filedVal := elem.FieldByName(filed)
	if !filedVal.IsValid() {
		return ErrFieldNotExist
	}

	if filedVal.Type().Name() != reflect.TypeOf(val).Name() {
		return ErrValueTypeNotMatch
	}

	if filedVal.CanSet() {
		filedVal.Set(reflect.ValueOf(val))
	}
	return nil

}
