package main

import "fmt"

func main() {
	var voteMap map[string]int = map[string]int{"gaurav": 31, "Nidhi": 41}
	var voteMap1 = map[string]int{"gaurav": 31, "Nidhi": 41}
	voteMap2 := map[string]int{"gaurav": 31, "Nidhi": 41}
	var voteMap3 map[string]int = make(map[string]int)
	var voteMap4 map[string]int
	fmt.Println(voteMap, voteMap1, voteMap2, voteMap3, voteMap4)
	//Append test on nil map
	voteMap3["Angad"] = 14
	fmt.Println(voteMap3)
	val, isKeyPresent := voteMap3["test"]
	fmt.Println("Non existeng key", val, isKeyPresent)
	fmt.Println("Before delete", voteMap)
	delete(voteMap, "gaurav")
	fmt.Println("After delete", voteMap)
}
