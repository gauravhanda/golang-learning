package main

import "fmt"

// Meters: Represents Meters by extending float64
type Meters float64

// Meters: Represents Meters by extending float64
type Centimeters float64

// Convert one type to other even if they are from same basetype
func (m Meters) ToCentimeters() Centimeters {
	return Centimeters(m * 1000)
}

func main() {
	var meters = Meters(1)
	var inCms = meters.ToCentimeters()
	fmt.Println(inCms)
}
