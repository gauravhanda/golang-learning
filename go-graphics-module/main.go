package main

import (
	"fmt"

	"github.com/gauravhanda/mathlib"
	"github.com/gauravhanda/mathlib/algebra"
	"github.com/gauravhanda/mathlib/calculus"
	"github.com/gauravhanda/mathlib/geometry"
	"github.com/gauravhanda/mathlib/trigonometry"
)

func main() {
	messageFromMathLib := mathlib.Intro()
	fmt.Println(messageFromMathLib)
	fmt.Println(algebra.Intro())
	fmt.Println(trigonometry.Intro())
	fmt.Println(calculus.Intro())
	fmt.Println(geometry.Intro())
}
