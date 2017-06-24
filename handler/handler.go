package handler

import (
	"net"
	"reflect"
)

var handlerMap = make(map[string]interface{})

func init() {
	handlerMap["testHandler"] = new(TestHandler)
}

func dispatcher(interfaceName string, methodName string, conn net.TCPConn) {
	reflect.ValueOf(handlerMap[interfaceName]).MethodByName(methodName).Call([]reflect.Value{})

}
