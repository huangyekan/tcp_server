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
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	server.Run()
}
