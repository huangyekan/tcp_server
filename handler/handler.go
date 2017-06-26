package handler

import (
	"reflect"
	"strings"
	"tcp_server/protol"
)

var HandlerMap = make(map[string]interface{})

func init() {
	HandlerMap["testHandler"] = new(TestHandler)
}

func Dispatcher(content *protol.Content) *protol.Response {
	interfaceName, methodName := parseContent(content)
	params := buildParams(content)
	result := reflect.ValueOf(HandlerMap[interfaceName]).MethodByName(methodName).Call(params)
	return result[0].Interface().(*protol.Response)

}

func parseContent(content *protol.Content) (interfaceName string, methodName string) {
	values := strings.Split(content.Method, ".")
	return values[0], values[1]
}

func buildParams(content *protol.Content) []reflect.Value {
	result := make([]reflect.Value, 1)
	result[0] = reflect.ValueOf(content.Params)
	return result
}
