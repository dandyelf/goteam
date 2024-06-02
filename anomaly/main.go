package main

import (
	an "anomaly/analyser"
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"time"
)

var (
	StdAnomalyCoefficient float64
	Distribution          []float64
	Mean, STDdev          float64
)

func init() {
	flag.Float64Var(&StdAnomalyCoefficient, "k", 0.0000005, "STD Anomaly coefficient")
	flag.Parse()
}

func main() {
	fmt.Println("STD Anomaly coefficient: ", StdAnomalyCoefficient)
	scanner := bufio.NewScanner(bufio.NewReader(os.Stdin))
	for scanner.Scan() {
		value, err := strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			fmt.Println("Error parsing input:", err)
			continue
		}
		err = an.AnomalyAnalise(value, StdAnomalyCoefficient)
		if err != nil {
			fmt.Println("Error parsing input:", err)
			continue
		}
		time.Sleep(time.Second / 50)
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading from stdin:", err)
		return
	}
	fmt.Println("Stream ended.")
}
