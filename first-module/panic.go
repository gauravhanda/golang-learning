package main

import "fmt"

func main() {
	var returnFirst = panicAndRecoverWithNamedReturnValue()
	fmt.Printf("Paniced method returned %#v\n", returnFirst)

	returnFirst = panicAndRecoverWithAnonymousFunction()
	fmt.Printf("Paniced method returned %#v\n", returnFirst)

}

// See the power of named result
func panicAndRecoverWithNamedReturnValue() (message string) {
	defer deferredMethodUsingPointer(&message)
	fmt.Println("Triggering Panic")
	panic("Blew up through panic method")
	return
}

func deferredMethodUsingPointer(returnMessage *string) {
	fmt.Println("deferredMethodUsingPointer ==> ", recover())
	*returnMessage = "deferredMethodUsingPointer: Recovered from Panic"
}

// See the power of named result
func panicAndRecoverWithAnonymousFunction() (message string) {
	defer func() {
		fmt.Println("panicAndRecoverWithAnonymousFunction ", recover())
		message = "Failed in Anonymous function"
	}()
	fmt.Println("Triggering Panic")
	panic("Blew up through panic method")
	return
}
