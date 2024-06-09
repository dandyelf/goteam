package main

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"log"
	"os"
	pro "reporter/prot"
	"strconv"
	"strings"
	"time"

	"github.com/golang/protobuf/ptypes/timestamp"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter num of values: ")
	text, _ := reader.ReadString('\n')

	conn, err := grpc.NewClient(":3333", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()
	cli := pro.NewRepServClient(conn)
	fmt.Println(text)
	data, err := cli.GetData(context.Background(), nil)
	if err != nil {
		log.Fatalln(err)
	}
	num, err := strconv.Atoi(strings.Trim(text, "\n"))
	if err != nil {
		log.Fatalln(err)
	}
	for i := 0; i < num; i++ {
		g, err := data.Recv()
		time.Sleep(time.Second / 2)
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Println(err)
		}
		fmt.Println(g.GetSessionId(), g.GetFrequency(), ConvertToUTC(g.GetTimestamp()))
	}
}

func ConvertToUTC(ts *timestamp.Timestamp) time.Time {
	loc, _ := time.LoadLocation("Europe/Moscow")
	t := time.Unix(ts.GetSeconds(), int64(ts.GetNanos())).In(loc)
	return t
}
