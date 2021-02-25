package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	var currentTime time.Time = time.Now()
	fmt.Println("hello ", currentTime.Year())
	name := "Gaurav Handa"
	replacer := strings.NewReplacer("Handa", "Sharma")
	updatedString := replacer.Replace(name)
	fmt.Println(updatedString)
	printNames("gaurav", "handa", false)
}

func printNames(firstName string, lastName string, euFormat bool) {
	fmt.Println(firstName, ",", lastName)
}
