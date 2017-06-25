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
		if err != nil {
			if n == 0 {
				break;
			} else {
				log.Println("读取数据异常", err.Error())
			}
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
	response, err := handler.Dispatcher(&message.Content)
	if err != nil {
		response = protol.NETWORK_ERROR
	}
	res, err := json.Marshal(response)
	if err != nil {
		log.Println("返回参数序列话错误", err)
		response = protol.NETWORK_ERROR
	}
	ConnMap[message.Header.Token].Write(res)
}

func doRegister(conn *net.TCPConn, message *protol.Message) {
	log.Println(" register ", "ip : ", conn.RemoteAddr().String(), "token : ", message.Header.Token)
	ConnMap[message.Header.Token] = conn
}

func ping(conn *net.TCPConn) {
	conn.Write([]byte("success..."))
}

func parseMessage(message string) *protol.Message {
	result := new(protol.Message)
	err := json.Unmarshal([]byte(message), &result)
	if err != nil {
		log.Println("json解析异常", err.Error())
	}
	return result
}
