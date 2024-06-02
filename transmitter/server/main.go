package main

import (
	"fmt"
	"log"
	"math/rand"
	"net"
	pro "reporter/prot"
	"time"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/timestamppb"
)

const (
	minMean = -10
	maxMean = 10
	stdLow  = 0.3
	stdHigh = 1.5
)

type reportServer struct {
	pro.UnimplementedRepServServer
}

func (s reportServer) GetData(empty *empty.Empty, stream pro.RepServ_GetDataServer) error {
	const StreamMaxVals = 100000000000000000
	uuidStr := uuid.NewString()
	mean, stdDev := MeanStdDevGenerator()
	log.Println(uuidStr, mean, stdDev, ConvertToUTC(timestamppb.Now()))
	for i := 0; i < StreamMaxVals; i++ {
		frequency := Frequency(mean, stdDev)
		timestamp := timestamppb.Now()
		response := pro.Point{
			SessionId: uuidStr,
			Frequency: frequency,
			Timestamp: timestamp,
		}
		if err := stream.Send(&response); err != nil {
			log.Println("stream stop")
			return err
		}
		time.Sleep(time.Second / 50)
	}
	return nil
}

func main() {
	lis, err := net.Listen("tcp", ":3333")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)
	pro.RegisterRepServServer(grpcServer, reportServer{})
	fmt.Println("Server start")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal(err)
	}
	log.Println("connecting")
}

func MeanStdDevGenerator() (mean float64, stdDev float64) {
	mean = rand.Float64()*(maxMean-minMean) + minMean // Random mean between -10 and 10
	stdDev = rand.Float64()*(stdHigh-stdLow) + stdLow // Random standard deviation between 0.3 and 1.5
	return
}

func Frequency(mean float64, stdDev float64) float64 {
	return rand.NormFloat64()*stdDev + mean
}

func ConvertToUTC(ts *timestamp.Timestamp) time.Time {
	t := time.Unix(ts.GetSeconds(), int64(ts.GetNanos())).UTC()
	return t
}
