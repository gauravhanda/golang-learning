package main

import "fmt"

func main() {
	var birthDay Date
	birthDay.SetDay(16)
	birthDay.SetYear(1980)
	birthDay.SetMonth(2)
	fmt.Println(birthDay)
	fmt.Println(birthDay.Year(), birthDay.Month(), birthDay.Day())

	var event Event = Event{Title: "My Birthday"}
	event.SetDay(16)
	event.SetMonth(2)
	event.SetYear(1980)
	fmt.Println(event)
	fmt.Println(event.Year(), event.Month(), event.Day())

}
