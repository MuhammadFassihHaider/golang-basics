package main

import "fmt"

func looping() {
	fmt.Println("looping")
	/*
		Keywords:
		for, continue, break
	*/

	for i := 0; i < 5; i++ {
		fmt.Println("loop", i)
	}

	/*
		We have have only one statement for each section
		of the 'for' loop. 'i++' counts as one statement
		in itself, so we cannot use it.
	*/
	for i, j := 0, 0; i < 5; i, j = i+1, j+2 {
		fmt.Println(i, j)
	}

	/*
		The difference is scope
	*/
	i := 0
	for ; i < 5; i++ {
		fmt.Println("loop", i)
	}

	/*
		WHILE_LOOP
	*/
	for i < 5 {
		fmt.Println("loop", i)
		i++
	}

	// OR

	for i < 5 {
		fmt.Println("loop", i)
		i++
	}

	/*
		INFINITE_LOOP
	*/
	for {
		fmt.Println("loop", i)
		break
	}

	/*
		Labelled code blocks
	*/
CODE_BLOCK:
	for i := 0; i <= 3; i++ {
		for j := 0; j <= 3; j++ {
			fmt.Println("CODE_BLOCK: ", i*j)
			if i*j >= 3 {
				break CODE_BLOCK

			}
		}
	}

	/*
		Looping over slices, arrays and maps
	*/
	/*
		Slices & arrays
	*/
	myArray := []int{1, 2, 3}
	for k, v := range myArray {
		fmt.Println(k, v)
	}

	/*
		Maps
	*/
	myMap := map[string]int{
		"English":     10,
		"Mathematics": 20,
		"Computers":   30,
	}
	for k, v := range myMap {
		fmt.Println(k, v)
	}

	/*
		Strings
	*/
	myString := "Hello Golang!"
	for k, v := range myString {
		fmt.Println(string(v), k, v)
	}
}
