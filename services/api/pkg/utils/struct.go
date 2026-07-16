/**

Filename: 		struct.go
Author: 		gengming - gengming.zb@ccbft.com
Description:	loongrpc struct tool logic
Create:			2022-07-02 11:15:05
Last Modified:	2022-07-02 14:30:27

*/

package utils

import (
	"reflect"
	"time"
)

func StructToCameCaseMap(data interface{}) map[string]interface{} {
	m := map[string]interface{}{}
	t := reflect.TypeOf(data)
	v := reflect.ValueOf(data)
	for i := 0; i < t.NumField(); i++ {
		isUnexported := t.Field(i).PkgPath != ""
		if isUnexported {
			continue
		}
		k := t.Field(i).Tag.Get("json")
		if k != "" {
			if k == "-" {
				//json遇到-标签则忽略
				continue
			}
			m[CamelCase(k)], _ = InterfaceToCamelCase(v.Field(i).Interface())
		} else if v.Field(i).Kind() == reflect.Struct {
			switch v.Field(i).Interface().(type) {
			case time.Time:
				m[FirstToLower(t.Field(i).Name)] = v.Field(i).Interface()
			default:
				for k, v := range StructToCameCaseMap(v.Field(i).Interface()) {
					m[k] = v
				}
			}
		} else {
			m[FirstToLower(t.Field(i).Name)], _ = InterfaceToCamelCase(v.Field(i).Interface())
		}
	}
	return m
}
