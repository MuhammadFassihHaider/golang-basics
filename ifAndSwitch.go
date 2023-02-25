package main

import "fmt"

func ifInGolang() {
	/*
		Simplest if condition
	*/
	if true {
		fmt.Println("I am true")
	}

	/*
		We can also have initializer if condition. The variables
		from the initializer are scoped to the if block.
	*/
	myMarks := map[string]int{
		"Science": 85,
		"Maths":   100,
	}
	if science, ok := myMarks["Science"]; ok {
		fmt.Println(science, ok)
	}

	/*
		Logical operators:
		>, <, ==, >=, <=, !=, !, &&, ||
	*/

	/*
		Short circuiting.

		Golang runs the condition lazily. When it finds that
		num1 is greater than 10, it enters the if block and does
		not computer to check if the other conditions are true or
		not. This is because, it will have to enter the if block and
		any other condition is irrelavant.

		For the second condition, it does calculate getTrueResult() and
		we see the print statement. This is because, the first case
		is false.
	*/
	const (
		num1 uint = 50
		num2 uint = 60
	)

	if num1 > 10 || getTrueResult() || num2 > 10 {
		fmt.Println("condition 1")
	}
	if num1 < 10 || getTrueResult() || num2 > 10 {
		fmt.Println("condition 2")
	}
}

func switchInGolang() {
	/*
		Simplest switch case
	*/
	switch 10 {
	case 1:
		fmt.Println("1")
	case 2:
		fmt.Println("2")
	default:
		fmt.Println("default")
	}

	/*
		We cannot have multiple cases for a single block like in other languages,
		however, we can have multiple tests in the same case as such
	*/
	switch 10 {
	case 1, 3, 5, 7, 9, 11:
		fmt.Println("1")
	case 2, 4, 6, 8, 10:
		fmt.Println("2")
	default:
		fmt.Println("default")
	}

	num := 10
	switch cond := num % 2; cond { // simply n % 2
	case 1:
		fmt.Println("1")
	case 0:
		fmt.Println("2")
	default:
		fmt.Println("default")
	}

	/*
		We can also have a switch where we put no condition
		infront of it
	*/
	num2 := 10
	switch { // simply n % 2
	case num2%2 == 1:
		fmt.Println("1")
	case num2%2 == 0:
		fmt.Println("2")
	default:
		fmt.Println("default")
	}

	/*
		There is no need for adding 'break' because it is implied by
		default. But 'break' keyword does exist in Goland and we can
		use it if we want to exit a case
	*/

	/*
		There is another keyword in context of switch called 'fallthrough'.
		Fallthrough is weird in its behaviour. If we use fallthrough, the
		condition for the case will not be evaluated and the statement in the
		next case will be run no matter what. It cannot be used conditionally
		as well.
		https://stackoverflow.com/questions/45268681/golangs-fallthrough-seems-unexpected
		The default behaviour in other languages is to fallthrough
	*/
	num3 := 11
	switch { // simply n % 2
	case num3%2 == 1:
		fmt.Println("1")
		fallthrough
	case num3%2 == 0:
		fmt.Println("fallthrough")
	default:
		fmt.Println("default")
	}
}

func typeSwitch() {
	var i interface{} = 5
	switch i.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	case float32:
		fmt.Println("float")
	}
}

func ifAndSwitch() {
	fmt.Println("ifAndSwitch")
	ifInGolang()
	switchInGolang()
	typeSwitch()
}

func getTrueResult() bool {
	fmt.Println("getTrueResult")
	return true
}
