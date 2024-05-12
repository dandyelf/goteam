package main

import "fmt"

func main() {
	err := Hallo()
	if err != nil {
		fmt.Println("Hallo, ERROR")
	}
}

func Hallo() error {
	var err error
	fmt.Println("Hallo world!")
	return err
}
