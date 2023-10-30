package function

import (
	"reflect"
	"strings"
)

// StructToMap 使用反射实现，完美地兼容了json标签的处理：修改参数处理
func StructToMap(st any) map[string]interface{} {
	m := make(map[string]interface{})
	val := reflect.ValueOf(st)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}
	if val.Kind() != reflect.Struct {
		return m
	}

	// 以下值忽略
	ignore := NewSliceContains[string]([]string{"-", "id", "created_at", "updated_at", "delete_at"})

	relType := val.Type()
	for i := 0; i < relType.NumField(); i++ {
		// 嵌套结构体递归
		if val.Field(i).Kind() == reflect.Struct {
			mapVal := StructToMap(val.Field(i).Interface())
			m = MapMerge[any](mapVal, m)
			continue
		}
		name := relType.Field(i).Name
		tag := relType.Field(i).Tag.Get("json")
		var tagList []string
		if tag != "" {
			tagList = strings.Split(tag, ",")
			name = tagList[0]
		}

		if ignore.ContainsSlice(tagList) {
			continue
		}

		if val.Field(i).IsZero() {
			m[name] = nil
		} else {
			m[name] = val.Field(i).Interface()
		}
	}
	return m
}
