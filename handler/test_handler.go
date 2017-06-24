package handler

import "net"

type TestHandler struct {

}

func (th *TestHandler) test(conn net.TCPConn)  {
	conn.Write([]byte("tesg"))
}

