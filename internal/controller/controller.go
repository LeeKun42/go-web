package controller

import (
	"github.com/gogf/gf/v2/net/ghttp"
)

type ResponseDataStruct struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type response struct {
	request *ghttp.Request
	data    *ResponseDataStruct
}

var res response

func Response(req *ghttp.Request) *response {
	res := &response{request: req, data: &ResponseDataStruct{Code: 0, Message: "ok", Data: nil}}
	return res
}

func (res *response) success(data interface{}) {
	res.data.Data = data
	res.request.Response.WriteJsonExit(res.data)
}

func (res *response) error(message string) {
	res.data.Code = 10000
	res.data.Message = message
	res.request.Response.WriteJsonExit(res.data)
}
