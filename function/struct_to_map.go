package function

import (
	"reflect"
	"strings"
)

// StructToMap 使用反射实现，完美地兼容了json标签的处理
func StructToMap(st any) map[string]interface{} {
	m := make(map[string]interface{})
	val := reflect.ValueOf(st)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}
	if val.Kind() != reflect.Struct {
		return m
	}

	relType := val.Type()
	for i := 0; i < relType.NumField(); i++ {
		name := relType.Field(i).Name
		tag := relType.Field(i).Tag.Get("json")
		if ValueInSlice[string]([]string{"-", "created_at", "updated_at", "delete_at"}, tag) {
			continue
		}
		if tag != "" {
			index := strings.Index(tag, ",")
			if index == -1 {
				name = tag
			} else {
				name = tag[:index]
			}
		}
		if val.Field(i).IsZero() {
			m[name] = nil
		} else {
			m[name] = val.Field(i).Interface()
		}
	}
	return m
}
