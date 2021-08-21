package main

import (
	"io"
	"log"
	"net/http"
	"net/rpc"
	"net/rpc/jsonrpc"
)

type HelloService struct{}

func (h *HelloService) Hello(request string, response *string) error {
	*response = "Hello " + request
	return nil
}

func main() {
	err := rpc.RegisterName("HelloService", &HelloService{})
	if err != nil {
		log.Fatal(err)
	}
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		var conn io.ReadWriteCloser = struct {
			io.Writer
			io.ReadCloser
		}{
			ReadCloser: r.Body,
			Writer:     w,
		}
		err = rpc.ServeRequest(jsonrpc.NewServerCodec(conn))
		if err != nil {
			log.Println(err)
		}
	})
	http.ListenAndServe(":1234", nil)
}
