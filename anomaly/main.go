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
	AnomalyCoefficient float64
	Distribution       []float64
	Mean, STDdev       float64
)

func init() {
	flag.Float64Var(&AnomalyCoefficient, "k", 0.0, "STD Anomaly coefficient")
	flag.Parse()
}

func main() {
	fmt.Println("STD Anomaly coefficient: ", AnomalyCoefficient)
	scanner := bufio.NewScanner(bufio.NewReader(os.Stdin))
	for scanner.Scan() {
		value, err := strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			fmt.Println("Error parsing input:", err)
			continue
		}
		Distribution = append(Distribution, value)
		Mean, STDdev, err = an.MeanStdDevCalc(Distribution)
		if err != nil {
			fmt.Println("Error calculating mean and STDdev:", err)
			continue
		}
		// Вычисляем скользящее mean и STDdev
		if len(Distribution)%6 == 0 {
			fmt.Println(len(Distribution), "value: ", value)
			fmt.Println("Mean: ", Mean, "STDdev: ", STDdev)
		}
		if len(Distribution)%6 == 100 {
			break
		}
		time.Sleep(time.Second / 20)
	}

	for scanner.Scan() {
		value, err := strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			fmt.Println("Error parsing input:", err)
			continue
		}
		err = an.AnomalyAnalise(value)

	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading from stdin:", err)
		return
	}
	fmt.Println("Stream ended.")
}
