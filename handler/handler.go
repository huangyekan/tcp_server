package handler

import (
	"reflect"
	"strings"
	"encoding/json"
	"log"
	"tcp_server/protol"
)

var HandlerMap = make(map[string]interface{})

func init() {
	HandlerMap["testHandler"] = new(TestHandler)
}

func Dispatcher(content *protol.Content) (*protol.Response, error){
	interfaceName, methodName := parseContent(content)
	params, err := buildParams(content)
	if err != nil {
		log.Println("参数解析出错", err)
		return nil, err
	}
	result := reflect.ValueOf(HandlerMap[interfaceName]).MethodByName(methodName).Call(params)
	return result[0].Interface().(*protol.Response), nil

}

func parseContent(content *protol.Content) (interfaceName string, methodName string) {
	values := strings.Split(content.Method, ",")
	return values[0], values[1]
}

func buildParams(content *protol.Content) ([]reflect.Value, error){
	if strings.Trim(content.Params, " " ) == "" {
		return nil, nil
	}
	var params map[string]interface{}
	err := json.Unmarshal([]byte(content.Params), &params)
	if err != nil {
		return nil, err
	}
	result := make([]reflect.Value, 1)
	result[0] = reflect.ValueOf(params)
	return result, nil
}
