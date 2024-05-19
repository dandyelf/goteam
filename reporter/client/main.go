package main

import (
	"context"
	"fmt"
	"log"
	pro "reporter/proc"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	serverAddr := "localhost:3333"
	conn, err := grpc.NewClient(serverAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()
	cli := pro.NewServClient(conn)
	var p pro.Point
	data, err := cli.GetData(context.Background(), &p)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(data.SessionId, data.Frequency, data.Timestamp)
}
