package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	seedValue := time.Now().Unix()
	rand.Seed(seedValue)
	randomNumber := rand.Intn(50)
	fmt.Println("Added random number", randomNumber)

	for i := 1; i <= 10; i++ {
		fmt.Println("Guess a number (Try ", i, ")")
		bufferIO := bufio.NewReader(os.Stdin)
		var userInput, err = bufferIO.ReadString('\n')
		if err == nil {
			userInput = strings.TrimSpace(userInput)
			var number int
			number, err := strconv.Atoi(userInput)

			if err == nil {
				_, err := checkNumber(number, randomNumber)
				if err != nil {
					fmt.Println(err)
				} else {
					fmt.Println("Match Found in retries ", i)
					break
				}
			} else {
				fmt.Println("Failed to read user input")
			}
		} else {
			log.Fatal("Failed to read user input with error")
		}
	}

}

func checkNumber(number int, randomNumber int) (bool, error) {
	if number < randomNumber {
		return false, errors.New("Number is too LOW")
	} else if number > randomNumber {
		return false, errors.New("Number is too HIGH")
	} else {
		return true, nil
	}

}
