package main

import (
	"fmt"
	"log"
	"math/rand"
	"net"
	pro "reporter/prot"
	"time"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type reportServer struct {
	pro.UnimplementedRepServServer
}

func main() {
	lis, err := net.Listen("tcp", ":3333")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	pro.RegisterRepServServer(grpcServer, reportServer{})
	fmt.Println("Server start")
	grpcServer.Serve(lis)
}

func (reportServer) GetData(empty *empty.Empty, stream pro.RepServ_GetDataServer) error {
	uuidStr := uuid.NewString()
	frequency := Frequency()
	timestamp := timestamppb.New(time.Now())
	log.Println(uuidStr, frequency, timestamp.String())

	response := pro.Point{
		SessionId: uuidStr,
		Frequency: frequency,
		Timestamp: timestamp,
	}
	if err := stream.Send(&response); err != nil {
		return err
	}
	return nil
}

func Frequency() float64 {

	const (
		minMean = -10
		maxMean = 10

		stdLow  = 0.3
		stdHigh = 1.5
	)

	mean := rand.Float64()*(maxMean-minMean) + minMean // Random mean between -10 and 10
	stdDev := rand.Float64()*(stdHigh-stdLow) + stdLow // Random standard deviation between 0.3 and 1.5

	return rand.NormFloat64()*stdDev + mean
}
