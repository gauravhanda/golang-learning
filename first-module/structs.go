// Demo of using structs to support mix types, arrays, slices and mapso only support one type
package main

import "fmt"

type myInt int32

// Subscriber : Defines the subscriber of our magazine
type Subscriber struct {
	name       string
	rate       float32
	isActive   bool
	postalCode myInt
}

func main() {
	fmt.Println("Learning structs to support mix types")
	//First style, inline declaration
	inlineDeclaration()

	// Reusing the type created on global scope of packages
	var gaurav Subscriber
	gaurav.name = "Gaurav"
	gaurav.rate = 4.98
	gaurav.isActive = true

	fmt.Println(gaurav)

	var nidhi Subscriber
	nidhi.name = "Nidhi"
	nidhi.rate = 4.98
	nidhi.isActive = false

	fmt.Println(nidhi)
}

// inlineDeclaration : defines inline declaration of structs. Non resuable style but possible
func inlineDeclaration() {
	var complexNumber struct {
		imaginary int32
		real      float32
	}
	complexNumber.imaginary = 6
	complexNumber.real = 7
	fmt.Println(complexNumber)

}
