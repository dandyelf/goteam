package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"net"
	pro "reporter/proc"
	"time"

	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type reportServer struct {
	pro.UnimplementedServServer
}

func main() {
	lis, err := net.Listen("tcp", ":3333")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	pro.RegisterServServer(grpcServer, reportServer{})
	fmt.Println("Server start")
	grpcServer.Serve(lis)
}

func (reportServer) GetData(context.Context, *pro.Point) (*pro.Point, error) {
	uuidStr := uuid.NewString()
	frequency := Frequency()
	var timestamp timestamppb.Timestamp
	log.Println(uuidStr, frequency, timestamp.String())

	return &pro.Point{
		SessionId: uuidStr,
		Frequency: frequency,
		Timestamp: &timestamp,
	}, nil

}

// func newServer() (s *reportServer) {

// 	return
// }

func Frequency() float64 {
	rand.Seed(time.Now().UnixNano())

	mean := rand.Float64()*20 - 10     // Random mean between -10 and 10
	stdDev := rand.Float64()*1.2 + 0.3 // Random standard deviation between 0.3 and 1.5

	return rand.NormFloat64()*stdDev + mean
}
