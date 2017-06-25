package main

import (
	"fmt"
	"tcp_server/server"
)

type Test struct {
	
}

func (t *Test) Test1() {
	fmt.Println("ccc")
}

func main() {
	server.Run()
}