// Demo of using structs to support mix types, arrays, slices and mapso only support one type
package main

import "fmt"

type myInt int32

type CString string

// Subscriber : Defines the subscriber of our magazine
type Subscriber struct {
	name       string
	rate       float32
	isActive   bool
	postalCode myInt
}
type Address struct {
	zipCode int8
	street  string
}

type Employee struct {
	id      int8
	address Address
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
	updateSubscriber(gaurav)
	updateSubscriberByPtr(&gaurav)

	fmt.Println(gaurav)

	var nidhi Subscriber
	nidhi.name = "Nidhi"
	nidhi.rate = 4.98
	nidhi.isActive = false
	updateSubscriberByDotPtr(&nidhi)

	fmt.Println(nidhi)

	defSubscriber := buildDefaultSubscriber()
	printInfo(defSubscriber)

	defaultEmployee := buildDefaultEmployee()
	fmt.Println(*defaultEmployee)

	var customString float64 = float64(6)
	fmt.Println(customString)
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

func printInfo(defSubscriber *Subscriber) {
	fmt.Println(*defSubscriber)
}

func updateSubscriber(subscriber Subscriber) {
	subscriber.postalCode = 55305
}

// resolving using the pointer
func updateSubscriberByPtr(subscriber *Subscriber) {
	(*subscriber).postalCode = 55305
}

// resolving the dot operator
func updateSubscriberByDotPtr(subscriber *Subscriber) {
	subscriber.postalCode = 55044
}

func buildDefaultSubscriber() *Subscriber {
	var defaultSusbcriber Subscriber
	defaultSusbcriber.name = "Subscriber"
	defaultSusbcriber.isActive = false
	defaultSusbcriber.rate = 10.99
	defaultSusbcriber.postalCode = 6004

	return &defaultSusbcriber
}

func buildDefaultEmployee() *Employee {
	// Define nested object
	address := Address{zipCode: 8, street: "Washington Street"}
	// define main object
	employee := Employee{id: 16, address: address}
	return &employee
}
