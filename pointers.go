package main

import "fmt"

func pointers() {
	fmt.Println("pointers")
	// creatingASimplePointer()
	// updadtingValuesUsingPointers()
	// usingArraysWithPointers()
	// usingStructsWithPointers()
	defaultPointerValue()
}

func creatingASimplePointer() {
	var a int = 42
	var b *int = &a // 'b' points to address of 'a'
	fmt.Println(a, b)
	/*
		'de-referencing' b to get the value at the address
		The '*' in *b is not the same as *int.
		'*int' says that the variable is of type pointer of data type 'int'
		'*b' is de-referencing to get its value
	*/
	fmt.Println(a, *b)
}

func updadtingValuesUsingPointers() {
	a := 42
	b := &a

	*b += 10 // add 10 to value at address b

	fmt.Println(a, *b)
}

func usingArraysWithPointers() {
	/*
		Pointer arithematics is not allowed in Golang. I.e
		We cannot add and subtract from pointer addresses
		to move along the addresses quickly to update values like
		in C++.
		We can use 'unsafe' package to do that if we really want to
	*/
	a := [3]int{1, 2, 3}
	b := &a

	(*b)[0] = 10
	(*b)[1] = 10
	(*b)[2] = 10

	fmt.Println(a, *b)
}

func usingStructsWithPointers() {
	var ms *myStruct
	ms = &myStruct{foo: 42}
	fmt.Println(ms)  // &{42}
	fmt.Println(*ms) // {42}

	// 42 || We need to add '()' because '.' operator has lower than '*'
	fmt.Println((*ms).foo)

	/*
		This is syntactic sugar. Compiler knows that we want to access the
		value at the address so it de-references the value for us
	*/
	fmt.Println(ms.foo)

	/*
		We can also use the 'new' built-in function to create a pointer
	*/
	a := new(myStruct)
	fmt.Println(*a)
}

type myStruct struct {
	foo uint64
}

func defaultPointerValue() {
	/*
		The default value for pointers is 'nil'
	*/
	var ms *myStruct
	fmt.Println(ms) // nil
}
