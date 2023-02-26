package main

import (
	"log"
	"net"

	c "github.com/prolifel/grpc-learn/config"
	pb "github.com/prolifel/grpc-learn/greet/proto"
	"google.golang.org/grpc"
)

type Server struct {
	pb.GreetServiceServer
}

func main() {
	config := c.GetConfig()

	listen, err := net.Listen("tcp", config.ServerHost+":"+config.ServerPort)
	if err != nil {
		log.Fatalf("💀 tcp failed to connect, msg: %s", err.Error())
	}

	log.Printf("server listening to %s:%s\n", config.ServerHost, config.ServerPort)

	server := grpc.NewServer()
	err = server.Serve(listen)
	if err != nil {
		log.Fatalf("💀 grpc server failed to listen, msg: %v", err.Error())
	}
}
