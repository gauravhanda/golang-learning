package main

import "fmt"

func main() {
	var slice []int
	slice = []int{1, 2, 34}
	fmt.Println(slice)
	slice = append(slice, 6, 7, 8)
	fmt.Println(slice)

	var makeSlice = make([]int, 5)
	fmt.Println(makeSlice)

	for index := range makeSlice {
		makeSlice[index] = index + 1
	}

	fmt.Println(makeSlice)

	var newSlice []int = makeSlice[:3]
	fmt.Println("New Slice", newSlice)
	newSlice[0] = 100
	fmt.Println("Before Apped", makeSlice)
	makeSlice = append(newSlice, 8, 9, 10)
	fmt.Println("Ater apppend", makeSlice)
	fmt.Println(makeSlice[2:5])

}
