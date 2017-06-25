package handler

import (
	"tcp_server/protol"
)

type TestHandler struct {

}

func (t *TestHandler) test(params map[string]interface{}) *protol.Response {
	result := make(map[string]interface{})
	result["name"] = "test"
	return &protol.Response{
		Msg:"ok",
		Code:100,
		Data:result,
	}
}

