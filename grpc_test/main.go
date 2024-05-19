package main

import (
	"fmt"
	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log"
	"math/rand"
	"net"
	ex00 "opp/proto"

	"time"
)

func main() {
	lis, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Fatal(err)
	}

	s := grpc.NewServer()
	ex00.RegisterTransmitterServer(s, server{})

	defer s.GracefulStop()
	fmt.Printf("start grpc server at port: %s\n", lis.Addr().String())
	if err := s.Serve(lis); err != nil {
		log.Fatal(err)
	}
}

type server struct {
	ex00.UnimplementedTransmitterServer
}

func (s server) StreamDataTransmit(empty *emptypb.Empty, stream ex00.Transmitter_StreamDataTransmitServer) error {
	uuidstr := uuid.NewString()
	mean, std := generateMeanAndStd()
	frequency := rand.NormFloat64()*mean + std
	log.Printf("UUID: %s, MEAN: %.2f, STD: %.2f", uuidstr, mean, std)

	resp := ex00.Response{
		SessionId: uuidstr,
		Frequency: frequency,
		Timestamp: timestamppb.New(time.Now()),
	}

	if err := stream.Send(&resp); err != nil {
		return err
	}

	return nil
}

func generateMeanAndStd() (mean float64, std float64) {
	const (
		minMean = -10
		maxMean = 10

		stdLow  = 0.3
		stdHigh = 1.5
	)

	mean = rand.Float64()*(maxMean-minMean) + minMean
	std = rand.Float64()*(stdHigh-stdLow) + stdLow

	return
}
