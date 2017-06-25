package protol

type Response struct {
	Msg string `msg`
	Code int `code`
	Data map[string]interface{} `data`
}

var NETWORK_ERROR *Response = &Response{
	Msg:"网络出错",
	Code:900,
	Data:nil,
}