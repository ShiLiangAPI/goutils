package res

import (
	"fmt"
	"net/http"
	"reflect"
)

type ApiError struct {
	Success bool `json:"success"`

	// Error code | 错误代码
	Code int `json:"code"`

	// Message | 提示信息
	Msg string `json:"msg"`

	// Data | 数据
	Data any `json:"data"`
}

func Error(msg any) *ApiError {
	t := reflect.TypeOf(msg).Kind()
	if t != reflect.String {
		msg = fmt.Sprintf("%s", msg)
	}
	return &ApiError{
		Success: false,
		Code:    http.StatusBadRequest,
		Msg:     msg.(string),
		Data:    map[string]any{},
	}
}

func BusinessError() *ApiError {
	return &ApiError{
		Success: false,
		Code:    http.StatusInternalServerError,
		Msg:     "系统繁忙",
		Data:    map[string]any{},
	}
}

func AuthError(msg string) *ApiError {
	return &ApiError{
		Success: false,
		Code:    http.StatusUnauthorized,
		Msg:     msg,
		Data:    map[string]any{},
	}
}
