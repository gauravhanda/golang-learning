package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	fileHandle, err := os.Open("data1.txt")

	if err != nil {
		log.Fatal("Failed to open File")
	}
	defer closeFile(fileHandle)
}

func closeFile(fileHandle *os.File) {
	fmt.Println("Closing file before exiting")
	err := fileHandle.Close()
	if err != nil {
		fmt.Println("Failed to close file")
	} else {
		fmt.Println("File closed succesfully")
	}

}
