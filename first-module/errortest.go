package main

import (
	"fmt"
	"log"
)

type MyError struct {
	errMessage string
}

func (m MyError) Error() string {
	return "MyCustomError Occurred"
}

func main() {
	result, err := divide(0, 10)
	if err != nil {
		log.Fatal(err, (err.(MyError)).errMessage)
	}

	fmt.Println(result)
}

func divide(divisor float32, dividend float32) (float32, error) {
	if divisor == 0 {
		return 0, MyError{errMessage: "Divison by Zero is not allowed"}
	}
	return dividend / divisor, nil
}
