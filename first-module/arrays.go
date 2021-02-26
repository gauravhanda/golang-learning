package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var intArray [3]int
	seed := time.Now().Unix()
	rand.Seed(seed)
	for i := 0; i < 3; i++ {

		intArray[i] = rand.Intn(60)
	}

	fmt.Println(intArray)
	intArray = [3]int{4, 6, 7}
	fmt.Println(intArray)
	for _, v := range intArray {
		fmt.Println(v)
	}
	fmt.Println("Lenght of an Array ", len(intArray))

	for index := range intArray {
		fmt.Println(index)
	}
}
