package handler

import (
	"tcp_server/protol"
)

type TestHandler struct {

}

func (t *TestHandler) Test(params map[string]interface{}) *protol.Response {
	result := make(map[string]interface{})
	result["name"] = "tcp_server"
	return &protol.Response{
		Msg:"ok",
		Code:100,
		Data:result,
	}
}

