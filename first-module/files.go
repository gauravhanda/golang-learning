package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("data.txt")
	if err != nil {
		log.Fatal("Failed to read file")
	}

	var numbers [3]float64
	scanner := bufio.NewScanner(file)

	for index := 0; scanner.Scan(); index++ {

		numbers[index], _ = convertToFloat(scanner.Text())
		fmt.Println(numbers[index])
	}
	fmt.Println("Arrays set to ", numbers)
	var sum, average float64

	for _, value := range numbers {
		sum += value
	}

	average = sum / float64(len(numbers))
	fmt.Println("Average = ", average)

	err = file.Close()
	if err != nil {
		fmt.Println("Failed to close the file")
	} else {
		fmt.Println("File closed successfully")
	}
}

func convertToFloat(value string) (float64, error) {
	return strconv.ParseFloat(value, 64)
}
