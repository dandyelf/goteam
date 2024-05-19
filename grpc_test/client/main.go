package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"io"
	"log"
	ex00 "opp/proto"
)

func main() {
	conn, err := grpc.NewClient(":8000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}

	client := ex00.NewTransmitterClient(conn)

	resp, err := client.StreamDataTransmit(context.Background(), nil)

	for {
		g, err := resp.Recv()
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Println(err)
		}

		fmt.Println(g)
	}

}
