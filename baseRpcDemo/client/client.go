package main

import (
	"fmt"
	"log"
	"net/rpc"
)

func main() {
	client, err := rpc.Dial("tcp", "localhost:1234")
	if err != nil {
		log.Fatal(err)
	}
	var resp string
	err = client.Call("HelloServices.Hello", "lance", &resp)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(resp)
}
