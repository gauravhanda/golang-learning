package main

import (
	"log"
	"os"
)

func main() {
	log.Println("Staring application")
	fileInfo, err := os.Stat("main.go1")
	var status bool
	if err != nil {
		status = true
	}

	if status == true {
		log.Println("File Found with size", fileInfo.Size())
	} else {
		log.Println(fileInfo.Size())
	}

}
