package main

import (
	"log"
	"net"
	"net/rpc"
)

type HelloServices struct{}

func (h *HelloServices) Hello(request string, response *string) error {
	*response = "hello " + request
	return nil
}

func main() {
	listner, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal(err)
	}
	err = rpc.RegisterName("HelloServices", &HelloServices{})
	if err != nil {
		log.Fatal(err)
	}
	conn, err := listner.Accept()
	if err != nil {
		log.Fatal(err)
	}
	rpc.ServeConn(conn)
}
