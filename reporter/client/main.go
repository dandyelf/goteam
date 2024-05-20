package main

import (
	"context"
	"fmt"
	"io"
	"log"
	pro "reporter/prot"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.NewClient(":3333", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()
	cli := pro.NewRepServClient(conn)

	data, err := cli.GetData(context.Background(), nil)
	if err != nil {
		log.Fatalln(err)
	}
	for {
		g, err := data.Recv()
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Println(err)
		}

		fmt.Println(g)
	}
}
