package utils

import (
	"github.com/gin-gonic/gin"
	"reflect"
	"strconv"
	"time"
)

func InterfaceToCamelCase(data interface{}) (interface{}, error) {
	if data == nil {
		return data, nil
	}
	t := reflect.TypeOf(data)
	v := reflect.ValueOf(data)
	_, ok := t.MethodByName("MarshalJSON")
	if ok {
		return data, nil
	}
	switch t.Kind() {
	case reflect.Map:
		m := map[string]interface{}{}
		if d, ok := data.(map[string]interface{}); ok {
			for k, v := range d {
				m[CamelCase(k)], _ = InterfaceToCamelCase(v)
			}
			return m, nil
		} else if d, ok := data.(gin.H); ok {
			for k, v := range d {
				m[CamelCase(k)], _ = InterfaceToCamelCase(v)
			}
			return m, nil
		}
	case reflect.Slice:
		length := v.Len()
		s := make([]interface{}, 0)
		for i := 0; i < length; i++ {
			x, _ := InterfaceToCamelCase(v.Index(i).Interface())
			s = append(s, x)
		}
		return s, nil
	case reflect.Struct:
		switch data.(type) {
		case time.Time:
			return data, nil
		default:
			return StructToCameCaseMap(data), nil
		}
	case reflect.Ptr:
		if v.IsNil() {
			return nil, nil
		}
		return InterfaceToCamelCase(v.Elem().Interface())
	case reflect.Int, reflect.Int8, reflect.Int16,
		reflect.Int32, reflect.Int64:
		return strconv.FormatInt(v.Int(), 10), nil
	case reflect.Uint, reflect.Uint8, reflect.Uint16,
		reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return strconv.FormatUint(v.Uint(), 10), nil
	}
	return data, nil
}
