package main

import "fmt"

func functions() {
	firstName := "Fassih"
	lastName := "Haider"

	functionWithCommonTypes(firstName, lastName)
	pointersAsArguments(&firstName)

	fmt.Println(firstName, lastName, "functions")

	variadicParameters("The sum is", 1, 2, 3, 4, 5)

	pointerResult := returningPointer(1, 2, 3, 4, 5)
	fmt.Println(pointerResult, *pointerResult, "functions")

	namedReturnType(1, 2, 3, 4, 5)

	result, err := multipleReturnTypes(6.0, 3.0)
	if err != nil {
		fmt.Println(err, "There was error in multipleReturnTypes")
	}
	fmt.Println(result, "multipleReturnTypes")

	anonymousFunctions("Hello Go!")

	methods()
}

// func functionWithCommonTypes(firstName, lastName string, age, height uint) | <- Invalid signature
func functionWithCommonTypes(firstName, lastName string) {
	fmt.Println(firstName, lastName, "functionWithCommonTypes")
}

func pointersAsArguments(firstName *string) {
	*firstName = "Muhammad"
}

func calculateSum(nums ...int) int {
	result := 0
	for _, v := range nums {
		result += v
	}
	return result
}

func variadicParameters(resultStatement string, nums ...int) {
	/*
		The variadic parameter has to be the last parameter
	*/
	result := calculateSum(nums...) // spreading values
	fmt.Println(resultStatement, result, "variadicParameters")
}

func returningPointer(nums ...int) *int {
	/*
		Golang recognizes that we are returning a pointer.
		So when the execution of the function ends, before the
		execution stack is destroyed, Golang promotes the value
		to the shared heap memory
	*/
	result := 0
	for _, v := range nums {
		result += v
	}
	return &result
}

func namedReturnType(nums ...int) (result int) {
	for _, v := range nums {
		result += v
	}
	fmt.Println(result, "namedReturnType")
	return
}

func multipleReturnTypes(a, b float64) (float64, error) {
	if b == 0.0 {
		return 0, fmt.Errorf("divide by 0 not allowed")
	}

	return a / b, nil
}

func anonymousFunctions(msg string) {
	func() {
		fmt.Println("Inside anonymous function")
	}()

	func(msg string) {
		fmt.Println(msg)
	}(msg)

	var functionInVariable func(msg string) = func(msg string) {
		fmt.Println(msg)
	}
	functionInVariable("I am a function stored in a variable")
}

func methods() {
	g := greeter{
		greeting: "Hello",
		name:     "Go!",
	}
	g.greet()
	g.changeName("Fassih")
	g.greet()

}

func (g greeter) greet() {
	/*
		This is a method
	*/
	fmt.Println(g.greeting, ",", g.name)
}

func (g *greeter) changeName(name string) {
	/*
		We can also have pointer to the struct
	*/
	g.name = name
}

type greeter struct {
	greeting string
	name     string
}
