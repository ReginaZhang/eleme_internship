package utils

import (
	"fmt"
	"reflect"
	"testing"
)

func TestMapType(t *testing.T) {
	_, ok := interface{}(make(map[string]interface{})).(Map)
	fmt.Println(ok)

	_, ok = interface{}(make(Map)).(map[string]interface{})
	fmt.Println(ok)

	fmt.Println(reflect.TypeOf(make(Map)).Kind())
}
