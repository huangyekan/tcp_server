package server

import (
	"net"
	"log"
	"encoding/json"
	"tcp_server/protol"
	"tcp_server/handler"
)

var ConnMap = make(map[string]*net.TCPConn)

var messageChannel = make(chan string, 16)

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
		go handleConnection(conn)
		defer conn.Close()
	}
}

func handleConnection(conn *net.TCPConn) {
	buf := make([]byte, 1024)
	tmpBuf := make([]byte, 0)
	go handleMessage(conn)
	for {
		n, err := conn.Read(buf)
		buf := append(buf[:n], tmpBuf...)
		log.Println(string(buf))
		if err != nil {
			log.Println("读取数据异常", err.Error())
			break;
		}
		messages, leftBuf := protol.Decode(buf, []string{})
		for _, message := range messages {
			log.Println("message : ", message)
			messageChannel <- message
		}
		tmpBuf = leftBuf
	}

}

func handleMessage(conn *net.TCPConn) {
	for {
		message := <-messageChannel
		msg := parseMessage(message)
		switch msg.Header.Type {
		case PING:
			ping(conn)
		case REQUEST:
			doRequest(msg)
		case REGISTER:
			doRegister(conn, msg)
		}
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
	log.Println(conn.RemoteAddr().String())
	conn.Write([]byte("aaa"))
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
	result := new(protol.Message)
	err := json.Unmarshal([]byte(message), &result)
	if err != nil {
		log.Println("json解析异常", err.Error())
	}
	return result
}
