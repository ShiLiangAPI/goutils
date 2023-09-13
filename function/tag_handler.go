package function

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"reflect"
	"strings"
)

// GetStructMsg 返回结构体中的msg参数
func GetStructMsg(err error, obj any) string {
	// 使用的时候，需要传obj的指针
	getObj := reflect.TypeOf(obj)
	// 将err接口断言为具体类型
	if errs, ok := err.(validator.ValidationErrors); ok {
		// 断言成功
		for _, e := range errs {
			// 循环每一个错误信息
			// 根据报错字段名，获取结构体的具体字段
			if f, exits := getObj.Elem().FieldByName(e.Field()); exits {
				msg := f.Tag.Get("msg")
				if msg != "" {
					if strings.Contains(msg, ";") {
						msgList := strings.Split(msg, ";")
						for _, msgValue := range msgList {
							keyValueList := strings.SplitN(msgValue, ":", 2)
							if strings.Contains(keyValueList[0], ",") {
								keyList := strings.SplitN(keyValueList[0], ",", 2)
								for _, key := range keyList {
									if key == e.Tag() {
										return keyValueList[1]
									}
								}
							} else if keyValueList[0] == e.Tag() {
								return keyValueList[1]
							}
						}
					} else {
						if strings.Contains(msg, ":") {
							keyValueList := strings.SplitN(msg, ":", 2)
							return keyValueList[1]
						}
						return msg
					}
				}
				gorm := f.Tag.Get("gorm")
				if gorm != "" {
					gormMsgList := strings.Split(gorm, ";")
					for _, gormMsg := range gormMsgList {
						keyValueList := strings.SplitN(gormMsg, ":", 2)
						if keyValueList[0] == "comment" {
							return keyValueList[1] + "不能为空"
						}
					}
				}
			}
		}
	}

	return err.Error()
}

//// GetValidComment 返回结构体中的comment参数
//func GetValidComment(err error, obj any) string {
//
//	getObj := reflect.TypeOf(obj)
//	if f, exits := getObj.Elem().FieldByName(err.Field()); exits {
//		tagValue := f.Tag.Get("gorm")
//		reg := regexp.MustCompile(`comment:'(.*?)'`)
//		matchArr := reg.FindStringSubmatch(tagValue)
//		return matchArr[len(matchArr)-1]
//	}
//
//	return ""
//}

// StructToFilterSlice 使用反射实现，完美地兼容了filter标签的处理
func StructToFilterSlice(st any) []string {
	var s []string
	val := reflect.ValueOf(st)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}
	if val.Kind() != reflect.Struct {
		return s
	}

	relType := val.Type()
	for i := 0; i < relType.NumField(); i++ {
		// 判断值是否为空
		if val.Field(i).IsZero() {
			continue
		}
		value := val.Field(i).Interface()
		// 获取标签对应的字段名称
		tag := relType.Field(i).Tag
		name := tag.Get("field")
		if name == "" {
			name = relType.Field(i).Tag.Get("form")
		}
		if name == "" {
			continue
		}
		index := strings.Index(name, ",")
		if index != -1 {
			name = name[:index]
		}
		// 获取筛选条件
		filter := strings.ToUpper(tag.Get("filter"))
		switch strings.ToUpper(filter) {
		case "":
			name += fmt.Sprintf(" = %v", value)
		case "IN":
			name += fmt.Sprintf(" IN (%v)", value.(string))
		case "LIKE":
			name += fmt.Sprintf(" LIKE '%%%v%%'", value)
		case "LLIKE":
			name += fmt.Sprintf(" LIKE '%%%v'", value)
		case "RLIKE":
			name += fmt.Sprintf(" LIKE '%v%%'", value)
		case "-":
			continue
		default:
			name += fmt.Sprintf(" %s %v", filter, value)
		}
		s = append(s, name)
	}
	return s
}
