package main

import (
	"log"

	c "github.com/prolifel/grpc-learn/config"
	"google.golang.org/grpc"
)

func main() {
	config := c.GetConfig()

	conn, err := grpc.Dial(config.ServerHost + ":" + config.ServerPort)
	if err != nil {
		log.Fatalf("💀 cannot dial to server, msg: %v", err.Error())
	}

	defer conn.Close()
}
