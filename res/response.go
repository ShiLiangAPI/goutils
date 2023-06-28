package res

import (
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
)

type Body struct {
	Success bool        `json:"success"`
	Code    int         `json:"code"`
	Msg     string      `json:"msg"`
	Data    interface{} `json:"data,omitempty"`
}

func Response(w http.ResponseWriter, resp interface{}, err error) {
	var body Body
	if err != nil {
		body.Success = false
		body.Code = -1
		body.Msg = err.Error()
	} else {
		body.Success = true
		body.Code = 200
		body.Msg = "请求成功"
		body.Data = resp
	}
	httpx.OkJson(w, body)
}
