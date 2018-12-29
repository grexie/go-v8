package v8

// #include "v8_c_bridge.h"
// #cgo CXXFLAGS: -I${SRCDIR} -I${SRCDIR}/include -g3 -fno-rtti -fpic -std=c++11
// #cgo LDFLAGS: -pthread -L${SRCDIR}/libv8 -lv8_base -lv8_init -lv8_initializers -lv8_libbase -lv8_libplatform -lv8_libsampler -lv8_nosnapshot
import "C"

import (
	"fmt"
	"reflect"
)

func (v *Value) Unmarshal(t reflect.Type) *reflect.Value {
	if t == valueType {
		v := reflect.ValueOf(v)
		return &v
	}

	switch t.Kind() {
	case reflect.Bool:
		v := reflect.ValueOf(v.Bool()).Convert(t)
		return &v
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		v := reflect.ValueOf(v.Int64()).Convert(t)
		return &v
	case reflect.Float32, reflect.Float64:
		v := reflect.ValueOf(v.Float64()).Convert(t)
		return &v
	case reflect.Array, reflect.Slice:
	case reflect.Func:
	case reflect.Ptr, reflect.Interface:
	case reflect.Map:
	case reflect.String:
		v := reflect.ValueOf(v.String()).Convert(t)
		return &v
	case reflect.Struct:
	}

	panic(fmt.Sprintf("unsupported kind: %v", t.Kind()))
}
