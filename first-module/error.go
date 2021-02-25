package main

import (
	"errors"
	"fmt"
)

func main() {
	value, test := numberChecker(1)
	fmt.Println(value, " -> ", test)
}

func numberChecker(x int) (string, error) {
	err := errors.New("Its done")
	var status string
	if x%2 == 0 {
		err = errors.New("Its even")
		status = "even"
	} else if x%2 != 0 {
		err = errors.New("Its odd")
		status = "odd"
	}

	return status, err
}
