package main

import (
	an "anomaly/analyser"
	"fmt"
	"os"
	"strconv"
)

func main() {
	args := os.Args[1:] // Получаем аргументы командной строки, начиная с первого

	numbers := make([]float64, 0) // Создаем пустой массив float64

	for _, arg := range args {
		num, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Printf("Input err %s value: %v\n", arg, err)
			continue
		}
		numbers = append(numbers, num)
	}
	fmt.Println(len(numbers))
	fmt.Println(an.AnomalyAnalyser(numbers))
}
