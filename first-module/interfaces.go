package main

import "fmt"

func main() {
	// Using the main type
	base64Encoder := MyBase64{message: "this is sample text"}
	var pointer = &base64Encoder
	fmt.Println(pointer.Encode())
	fmt.Println(pointer.Decode())

	var encoderInterface Encoding = &base64Encoder
	fmt.Println(encoderInterface.Decode())
	fmt.Println(encoderInterface.Encode())

}
