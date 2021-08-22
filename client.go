package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"time"

	pb "demoGRPC/services"
	"google.golang.org/grpc"
)

const (
	address     = "localhost:50051"
	defaultName = "world"
)

func init() {
	log.SetFlags(log.LstdFlags)
}
func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)

	// Contact the server and print out its response.
	name := defaultName
	if len(os.Args) > 1 {
		name = os.Args[1]
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.SayHello(ctx, &pb.HelloRequest{Name: name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetMessage())

	// 客户端流式发送.
	stream, err := c.StreamSayHello(context.Background())
	if err != nil {
		log.Println(err)
	}
	i := 0
	for {
		i++
		if err := stream.Send(&pb.HelloRequest{
			Name: name,
		}); err != nil {
			log.Println(err)
		}
		time.Sleep(time.Second)
		if i > 10 {
			break
		}
	}
	// 服务端流式发送.
	sendStream, err := c.StreamServer(context.Background(), &pb.HelloRequest{
		Name: name,
	})
	if err != nil {
		log.Println(err)
	}
	for {
		res, err := sendStream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Println(err)
			break
		}
		fmt.Println("receive from server message: " + res.Message)
	}
}
