package main

import (
	"fmt"
	"log"
	"tcp_server/server"
)

type Test struct {
}

func (t *Test) Test1() {
	fmt.Println("ccc")
}

func main() {
	//header := &protol.Header{
	//	Type:  "register",
	//	Token: "123456",
	//}
	//content := &protol.Content{
	//	Method: "testHandler.Test",
	//	Params: map[string]interface{}{
	//		"name": "132",
	//	},
	//}
	//message := &protol.Message{
	//	Header:  *header,
	//	Content: *content,
	//}
	//b ,_ := json.Marshal(message)
	//fmt.Println(string(b))
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	server.Run()
}
