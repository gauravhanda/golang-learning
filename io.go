package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	stream := bufio.NewReader(os.Stdin)
	var userInput, err = stream.ReadString('\n')
	if err == nil {
		log.Fatal("Call is successful !")
	} else {
		fmt.Println("Call failed")
	}
	fmt.Println(userInput, err)
}
