package main

import (
	"fmt"
)

var numberOne int
var numberTwo int
var total int
var do string

func main() {
	Calculator()
}

func Calculator() {

	fmt.Println("Type a numberOne: ")
	fmt.Scan(&numberOne)

	fmt.Println("Type a numberTwo: ")
	fmt.Scan(&numberTwo)

	fmt.Println("what to do with this + - * /: ")
	fmt.Scan(&do)

	if do == "+" {
		total = numberOne + numberTwo
	}

	if do == "-" {
		total = numberOne - numberTwo
	}

	if do == "/" {
		total = numberOne / numberTwo
	}

	if do == "*" {
		total = numberOne * numberTwo
	}

	fmt.Println("Your number is:", total)
}
