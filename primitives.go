package main

import (
	"fmt"
)

func primitives() {
	// Bool
	var myBoolTrue bool = true
	var myBoolFalse bool = false
	var myBoolDefault bool // false

	fmt.Println(myBoolTrue)
	fmt.Println(myBoolFalse)
	fmt.Println(myBoolDefault)

	// Numeric
	var (
		num1 int8  = 1
		num2 int16 = 2
		num3 int32 = 3
		num4 int64 = 4

		num5 uint8  = 5
		num6 uint16 = 6
		num7 uint32 = 7
	)
	fmt.Printf("%v %v %v %v \n", num1, num2, num3, num4)
	fmt.Printf("%v %v %v \n", num5, num6, num7)

	/*
		Floats are representation of decimals. It is very easily to have unexpected bugs.
		We should be careful when we are performing mathematical calculations
	*/
	var (
		pi               float32 = 3.14
		veryLargeNumber1 float64 = 2.4e84
	)
	fmt.Println(pi, veryLargeNumber1)

	// allowed arithematic (numeric): +, -, /, *, %
	// allowed arithematic (float): +, -, /, *

	// string
	var myName string = "Fassih Haider"
	fmt.Println(myName, myName[1], string(myName[1])) // Fassih Haider 97 a: Converts to bytes when we use string array
	var myNameAsBytes []byte = []byte(myName)         // uint8 array
	fmt.Println(myNameAsBytes)

	// runes
	// with single quotes
}
