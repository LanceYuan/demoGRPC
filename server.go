package main

import (
	"context"
	pb "demoGRPC/services"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	"time"
)

const (
	port = ":50051"
)

// server is used to implement helloworld.GreeterServer.
type server struct {
	pb.UnimplementedGreeterServer
}

// SayHello implements helloworld.GreeterServer
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Received: %v", in.GetName())
	return &pb.HelloReply{Message: "Hello " + in.GetName()}, nil
}

// 客户端流式发送.
func (s *server) StreamSayHello(stream pb.Greeter_StreamSayHelloServer) error {
	for {
		res, err := stream.Recv()
		if err != nil {
			log.Println(err)
			break
		}
		fmt.Println("receive from client message: " + res.Name)
	}
	return nil
}

// 服务端流式发送.
func (s *server) StreamServer(request *pb.HelloRequest, stream pb.Greeter_StreamServerServer) error {
	i := 0
	for {
		i++
		if err := stream.Send(&pb.HelloReply{
			Message: request.Name,
		}); err != nil {
			log.Println(err)
			break
		}
		if i > 10 {
			break
		}
		time.Sleep(time.Second)
	}
	return nil
}
func init() {
	log.SetFlags(log.LstdFlags)
}
func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
