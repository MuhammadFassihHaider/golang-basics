package main

import (
	"fmt"
	"math"
)

const i int8 = 32

func constants() {
	fmt.Println("CONSTANTS")

	/*
		const naming scheme is same as for var.
		Camel case for package scope variables
		Pascal case for global variables
	*/

	/*
		constant declaration cannot be used with
		functions. It gives compilation error
	*/
	var sinValueVar float64 = math.Sin(1.57)
	// const sinValueConst float64 = math.Sin(1.57)
	fmt.Println(sinValueVar)

	/*
		constants work with other primitive types
		like how we expect them to including
		int, string, bool, float32 etc
	*/

	/*
		if we have a constant in the package scope
		and another variable of same name in
		function scope, the function scope variable
		takes precedence
	*/
	var i int = 42
	println(i) // prints 42

	/*
		we can perform arithmetic between 'const'
		and 'var' variables. The result is a mutable
		value
	*/

	const num1 int = 5
	var num2 int = 10
	// const sum int = num1 + num2 // throws error
	var sum int = num1 + num2
	println(sum)

	/*
		we do not have to give type to constants
		as the are inferred. We also do not need
		to use ':=' syntax
	*/

	const num3 = 50
	// const num3 = 255
	// const num3 = 256
	println(num3)

	/*
		The example below works because the compiler
		replaces all the places where num3 is used
		with '57'. So when it tries to perform summation
		it takes 57 to be of uint8 type.
		This would not have worked if I had explicitly
		give type to the const declaration.
		Similary, if I increase the value of num3 to a value
		greater than 255 (max value of uint8), I get linting
		error.
		However, surprisingly, if I add 255 + 10 to uint8 value,
		it is an overflow. But this error is not give to me
		by the compiler or linter. Instead I just get 9 as result
	*/

	var num4 uint8 = 10
	println(num3 + num4)

	/*
		iota is a counter in golang
		We do not need to add iota againto be. This is because,
		if there is no initialization, the compiler sees a pattern
		and it uses the same formula
	*/

	const (
		a = iota // 0
		b        // 1
	)
	println(a, b)

	/*
		It can sometimes be dangerous to have the values starting from
		0 because of default values for numeric types.
		There are three things that we can do:
			a) we can add 1 to first value,
			b) use the first value as error value
			c) use '_'
	*/

	var undeclaredVariable int       // has a default value of 0
	println(undeclaredVariable == a) // true

	const (
		_ = iota
		c
		d
	)
	println(c, d)

	const (
		e = iota + 1
		f
	)
	println(e, f)

	const (
		errorValue = iota
		g
		h
	)

	println(g, h)

	/*
		shifting values
		need to understand it a bit better
			bitwise operations
	*/

	const (
		isAdmin = 1 << iota
		isHeadquarters
		canSeeFinancials

		isLocatedInAsia
		isLocatedInAfrica
		isLocatedInEurope
		isLocatedInNorthAmerica
		isLocatedInSouthAmerica
	)

	var roles byte = isAdmin | canSeeFinancials | isLocatedInEurope
	fmt.Printf("%b, %v\n", roles, roles)
	fmt.Printf("is Admin? %v \n", isAdmin&roles == isAdmin)
	fmt.Printf("is in Europe? %v \n", isLocatedInAsia&roles == isLocatedInAsia)
}
