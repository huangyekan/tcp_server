package server

import (
	"net"
	"log"
	"encoding/json"
)

var connMap = make(map[string]*net.TCPConn)

const (
	PING     = "ping"
	REQUEST  = "request"
	REGISTER = "register"
)

func Run() {
	lisener, err := net.ListenTCP("tcp",
		&net.TCPAddr{
			IP:   []byte("127.0.0.1"),
			Port: 8000,
			Zone: "",
		})
	if err != nil {
		log.Println("服务启动异常", err.Error())
	}
	for {
		conn, err := lisener.AcceptTCP();
		if err != nil {
			log.Println("连接失败", err.Error())
		}
		go handle(conn)
		defer conn.Close()
	}
}

func handle(conn *net.TCPConn) {
	buf := make([]byte, 1024)
	var message string
	for {
		n, err := conn.Read(buf)
		if err != nil {
			if n == 0 {
				break;
			} else {
				log.Println("读取数据异常", err.Error())
			}
		}
		message += string(buf[0:n])
	}
	msg := parseMessage(message)
	switch msg.msgType {
	case PING:
		ping(conn)
		break
	case REQUEST:
		request(conn, msg)
		break
	case REGISTER:
		register(conn, msg)
		break
	}

}

func request(conn *net.TCPConn, message *Message)  {
}

func register(conn *net.TCPConn, message *Message)  {
	log.Println(" register ", "ip : ", conn.RemoteAddr().String(), "token : ", message.msgContent.token)
	connMap[message.msgContent.token] = conn
}

func ping(conn *net.TCPConn)  {
	conn.Write([]byte("success"))
}

func parseMessage(message string) *Message {
	result := new(Message)
	err := json.Unmarshal([]byte(message), &result)
	if err != nil {
		log.Println("json解析异常", err.Error())
	}
	return result
}


type Message struct {
	msgType    string `type`
	msgContent MsgContent `content`
}

type MsgContent struct {
	token string `token`
	method string `method`
	params string `params`
}
