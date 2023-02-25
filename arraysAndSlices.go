package main

import "fmt"

func arraysAndSlices() {
	/*
		Arrays are VALUE type
		Structs are REFERENCE type
	*/
	/*
		Arrays: The elements in array are contigious in memory
	*/
	grades := [3]uint{95, 85, 89}
	subjects := [3]string{"Maths", "English", "Computer"}
	fmt.Println(grades)
	fmt.Println(subjects)

	/*
		In the above example, we are declaring the size of the
		array twice. One explicitly and the other by the number
		of elements that we add to the array during initialization.
		We can get rid of this by using the '[...]' syntax.
	*/
	grades2 := [...]uint{95, 85, 89}
	subjects2 := [...]string{"Maths", "English", "Computer"}
	fmt.Println(grades2)
	fmt.Println(subjects2)

	/*
		If we declare an array and do not add any values during
		initialization, the array is filled with default values
		for the type of array that we declared
	*/
	var grades3 [3]int
	fmt.Println(grades3)
	grades3[0] = 95
	fmt.Println(grades3)

	/*
		We can get the length of the array using the built in len function
	*/
	fmt.Println(len(grades3))

	/*
		We can have 2 dimentional arrays the following way
	*/
	myMatrix := [3][3]int{
		{0, 1, 2},
		{3, 4, 5},
		{6, 7, 8},
	}
	fmt.Println(myMatrix)

	/*
		Arrays are literally copied. They are not reference copies. For this
		reason, we need to be careful when we pass values to functions because
		arrays are passed by value and not by reference
	*/
	a := [...]int{1, 2, 3}
	b := a // 'b' points to a different location in memory than 'a
	b[0] = 0
	b[1] = 0
	b[2] = 0
	fmt.Println(a) // 1,2,3
	fmt.Println(b) // 0,0,0

	/*
		In slices, we do not need to specify the size before hand. Slices are passed
		by reference. Almost, everything that we can do with an array, we can also
		do with slices.

		Underlying the slice is an array. If we add an element to a slice and there
		is not memory block clear in front of the last element in the slice, Golang
		copies all the values of the slice to a new memory location.
		There is a  built-in function called cap(). This function gives us the
		projection of the size of the underlying array. For example
		We have a slice of len = 3. We add an element but there is not space in the
		next memory block. Now, Golang copies the values to a new memory location
		but instead of pasting the values in a new memory location with only 4
		spaces, the compiler pastes it at a places with more space than that so that
		it does not have to copy the values again when we add a new value to the slice.
		That size which Golang is expecting for our slice is the projecting size of the
		underlying array that we can get using the capacity(cap()) function
	*/
	c := []int{1, 2, 3}
	fmt.Println(c)
	fmt.Println(len(c))
	fmt.Println(cap(c))
	c = append(c, 100, 102, 103, 105)
	fmt.Println(c)
	fmt.Println(len(c)) // 7
	fmt.Println(cap(c)) // 8: Increase in capacity is random but generally is double the size of previous array
	/*
		Other methods of declaring slices
	*/
	d := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	e := d[:]  // make slice with all elements in 'd'
	f := d[3:] // start (including) the element on the 3 index i.e 4
	g := d[:6] // start from start and end till 5th index (excluding 6)
	h := d[3:6]
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
	fmt.Println(g)
	fmt.Println(h)
	g[2] = 100 // index 2 is 100 for all the values that have index 2
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
	fmt.Println(g)
	fmt.Println(h)

	/*
		We can make slice from a slice or an array using 'e := d[:]' syntax.
		Here 'd' can be a slice or an array
	*/

	/*
		We can also use the built in function 'make' to create a slice.
		make(typeof slice, initial len, cap)
	*/
	i := make([]int, 3, 100) // cap must be greater than len
	fmt.Println(i)
	fmt.Println(len(i))
	fmt.Println(cap(i))

	/*
		In order to append two slices together, we can use the append function
		with spread operator
	*/
	j := []int{1, 2, 3}
	k := []int{4, 5, 6}
	l := append(j, k...) // the second and rest of the args of append need to be of the type of slice passed in as first argument, hence we spread
	fmt.Printf("%v \n", l)

	/*
		removing first and last element in slice
	*/
	m := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	m = m[1:] // remove first element
	fmt.Println(m)
	m = m[:len(m)-1] // remove last element
	fmt.Println(m)
	m = append(m[:2], m[3:]...)
	fmt.Println(m)

	/*
		In order to make copies of slice, we will need to use loops
		or we can use the copy function
	*/
	n := []int{1, 2, 3}
	o := make([]int, len(n))
	copy(o, n)
	fmt.Println(n)
	fmt.Println(o)
	n = n[1:]
	fmt.Println(n) // 2,3
	fmt.Println(o) // 1,2,3
}
