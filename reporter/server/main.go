package main

import (
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	"reporter/proc"
)

func main() {

	lis, err := net.Listen("tcp", fmt.Sprintf(":3333"))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	proc.RegisterServServer(grpcServer)
	grpcServer.Serve(lis)
}
