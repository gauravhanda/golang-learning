package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	fmt.Println(args)
	fmt.Println(varArgs("Kgs", 1, 2, 7))
	slicedValue := []int{1, 3, 4}
	fmt.Println(varArgs("Kgs", slicedValue...))
}

func varArgs(unit string, numbers ...int) string {
	var sum int
	for _, v := range numbers {
		sum += v
	}
	return fmt.Sprintf("%d %s", sum, unit)
}
