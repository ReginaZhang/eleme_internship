package utils

import (
	"reflect"

	"github.com/pkg/errors"
	"time"
)

type Map map[string]interface{}

func NewMap(i interface{}) (Map, error) {
	switch m := i.(type) {
	case Map:
		return m, nil
	case map[string]interface{}:
		return Map(m), nil
	case nil:
		return nil, nil
	default:
		return nil, errors.Errorf("invalid type")
	}
}

func (m Map) Get(key string) interface{} {
	res, _ := m[key]
	return res
}

func (m Map) String(key string) string {
	i, _ := m.Get(key).(string)
	return i
}

func (m Map) StringSlice(key string) []string {
	i, _ := m.Get(key).([]string)
	return i
}

func (m Map) Float64(key string) float64 {
	i, _ := m.Get(key).(float64)
	return i
}

func EnsureStringKey(origin interface{}) interface{} {
	rt, rv := reflect.TypeOf(origin), reflect.ValueOf(origin)

	switch rt.Kind() {
	case reflect.Slice:
		var res []interface{}
		for i := 0; i < rv.Len(); i++ {
			res = append(res, EnsureStringKey(rv.Index(i).Interface()))
		}
		return res
	case reflect.Map:
		res := make(map[string]interface{})
		for _, k := range rv.MapKeys() {
			key, _ := k.Interface().(string)
			res[key] = EnsureStringKey(rv.MapIndex(k).Interface())
		}
		return res
	case reflect.Ptr:
		return EnsureStringKey(rv.Elem().Interface())
	default:
		return origin
	}
}

func FilterFields(obj interface{}, extraFields ...string) interface{} {
	rt, rv := reflect.TypeOf(obj), reflect.ValueOf(obj)

	blacklist := []string{"created_at", "updated_at", "drc_check_time", "-", ""}
	blacklist = append(blacklist, extraFields...)

	switch rt.Kind() {
	case reflect.Slice:
		var res []interface{}
		for i := 0; i < rv.Len(); i++ {
			res = append(res, FilterFields(rv.Index(i).Interface(), extraFields...))
		}
		return res
	case reflect.Struct:
		if _, ok := obj.(time.Time); ok {
			return obj
		}
		res := make(map[string]interface{})
		for i := 0; i < rt.NumField(); i++ {
			jsonTag := rt.Field(i).Tag.Get("json")
			if !contains(blacklist, jsonTag) {
				res[jsonTag] = FilterFields(rv.Field(i).Interface(), extraFields...)
			}
		}
		return res
	case reflect.Map:
		res := make(map[string]interface{})
		for _, k := range rv.MapKeys() {
			if !contains(blacklist, k.String()) {
				res[k.String()] = FilterFields(rv.MapIndex(k).Interface(), extraFields...)
			}
		}
		return res
	case reflect.Ptr:
		return FilterFields(rv.Elem().Interface(), extraFields...)
	default:
		return obj
	}

}

func contains(slice []string, s string) bool {
	for _, e := range slice {
		if e == s {
			return true
		}
	}
	return false
}
