package main
import (
	"fmt"
	"strconv"
)

func variables() {
	// three ways of declaring variables

	// method 1: good when I want to declare a variable and use it later
	var number1 int
	number1 = 42
	fmt.Println(number1)

	// method 2: when I want to have more control. Assume that i want number2 to be of type float
	var number2 int = 52
	fmt.Println(number2)

	// method 3: the method that I will normally use
	number3 := 62
	fmt.Println(number3)

	// ===============================================

	// variable block
	var (
		name   string  = "Fassih"
		email  string  = "fassih@gmail.com"
		age    int     = 24
		weight float32 = 65
	)
	fmt.Println(name, email, age, weight)

	// ===============================================

	// cannot re-declare the same variable in the same scope

	// variables at the closest scope are used. This is called as shadowing
	// for example, if there is a variable declared at the package level and another
	// is declared in the function, the variable in the function will take precedence

	// all variables have to be used in golang

	// ===============================================

	// three scopes in go: 1) package 2) global 3) block scope
	// variables starting with lower case are scoped to the package
	// variables starting with capital letter are globally available

	// type conversion in go needs to be explicit
	var number4 float32 = 32.598
	fmt.Printf("%v %T \n", number4, number4) // 32.598

	var number5 int = int(number4)
	fmt.Printf("%v %T \n", number5, number5) // 32

	// for string conversion, I will have to use a library strconv
	fmt.Println(strconv.Itoa(number5))
	
	// undeclared variables are given 0 values by default. For example
	// var myTest bool
	// myTest will have value of false
}
