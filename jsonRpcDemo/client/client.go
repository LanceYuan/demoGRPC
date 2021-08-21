package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc/jsonrpc"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:1234")
	if err != nil {
		log.Fatal(err)
	}
	client := jsonrpc.NewClient(conn)
	var resp string
	err = client.Call("HelloService.Hello", "lance", &resp)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(resp)
}
