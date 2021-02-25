package main

import (
	"fmt"
	"strconv"
)

func main() {
	PI := "3.1412"
	var value, err = strconv.ParseFloat(PI, 32)
	var value1, err = strconv.ParseFloat(PI, 32)
	fmt.Println(value)
	fmt.Println(err)

	fmt.Println(strconv.Itoa(97))
}
