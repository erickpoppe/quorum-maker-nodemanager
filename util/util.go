package util

import (
	"reflect"
	"strings"
	"strconv"
	"fmt"
	"github.com/magiconair/properties"
)

func TakeSliceArg(arg interface{}) (out []interface{}, ok bool) {
	slice, success := takeArg(arg, reflect.Slice)
	if !success {
		ok = false
		return
	}
	c := slice.Len()
	out = make([]interface{}, c)
	for i := 0; i < c; i++ {
		out[i] = slice.Index(i).Interface()
	}
	return out, true
}

func takeArg(arg interface{}, kind reflect.Kind) (val reflect.Value, ok bool) {
	val = reflect.ValueOf(arg)
	if val.Kind() == kind {
		ok = true
	}
	return
}

func HexStringtoInt64(hexVal string) (intVal int64) {
	hexVal = strings.TrimSuffix(hexVal, "\n")
	hexVal = strings.TrimPrefix(hexVal, "0x")
	intVal, err := strconv.ParseInt(hexVal, 16, 64)
	if err != nil {
		fmt.Println(err)
	}
	return intVal
}

func MustGetString(key string, filename *properties.Properties) (val string) {
	val = filename.MustGetString(key)
	val = strings.TrimSuffix(val, "\n")
	return val
}
