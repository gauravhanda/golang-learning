package main

import "fmt"

func main() {
	var radius int = 10
	var radiusPtr *int = &radius
	fmt.Println(radius, radiusPtr, &radius, *&radius, *radiusPtr)
}

func update(radius *int) {
	*radius = 9
}
