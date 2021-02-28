package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Starting goroutines")
	go loop("Gaurav")
	go loop("Nidhi")
	time.Sleep(5 * time.Second)
	fmt.Println("Main method ends")
}

func loop(name string) {
	for i := 1; i <= 5; i++ {
		fmt.Println(name, i)
	}
}


