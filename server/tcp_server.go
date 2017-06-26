package server

import (
	"net"
	"log"
	"encoding/json"
	"tcp_server/protol"
	"tcp_server/handler"
)

var ConnMap = make(map[string]*net.TCPConn)

const (
	PING     = "ping"
	REQUEST  = "request"
	REGISTER = "register"
)

func Run() {
	lisener, err := net.ListenTCP("tcp",
		&net.TCPAddr{
			IP:   net.ParseIP("127.0.0.1"),
			Port: 8000,
		})
	if err != nil {
		log.Println("服务启动异常", err.Error())
	}
	for {
		conn, err := lisener.AcceptTCP();
		log.Println("accept")
		if err != nil {
			log.Println("连接失败", err.Error())
			continue
		}
		go handle(conn)
		defer conn.Close()
	}
}

func handle(conn *net.TCPConn) {
	buf := make([]byte, 1024)
	var request string
	for {
		n, err := conn.Read(buf)
		log.Println(n)
		if n == 0 {
			break;
		}
		if err != nil {
			log.Println("读取数据异常", err.Error())
			break;
		}
		request += string(buf[0:n])
	}
	message := parseMessage(request)
	switch message.Header.Type {
	case PING:
		ping(conn)
		break
	case REQUEST:
		doRequest(message)
		break
	case REGISTER:
		doRegister(conn, message)
		break
	}

}

func doRequest(message *protol.Message) {
	response := handler.Dispatcher(&message.Content)
	res, _ := json.Marshal(response)
	log.Println("res ", string(res))
	conn := ConnMap[message.Header.Token]
	if conn == nil {
		log.Println("用户已断开")
		return
	}
	conn.Write(res)
}

func doRegister(conn *net.TCPConn, message *protol.Message) {
	log.Println(" register ", "ip : ", conn.RemoteAddr().String(), "token : ", message.Header.Token)
	ConnMap[message.Header.Token] = conn
}

func ping(conn *net.TCPConn) {
	log.Println(conn.RemoteAddr().String() + " ping success")
	conn.Write([]byte("success..."))
}

func parseMessage(message string) *protol.Message {
	log.Println(message)
	result := new(protol.Message)
	err := json.Unmarshal([]byte(message), &result)
	if err != nil {
		log.Println("json解析异常", err.Error())
	}
	return result
}
