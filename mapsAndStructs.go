package main

import (
	"fmt"
	"reflect"
)

func maps() {
	myGrades := map[string]int{
		"English":     80,
		"Urdu":        70,
		"Mathematics": 100,
		"Science":     92,
	}
	fmt.Println(myGrades)

	/*
		Adding and deleting a value from a map
	*/

	myGrades["Computers"] = 100 // adding
	delete(myGrades, "Urdu")    // deleting
	fmt.Println(myGrades)

	/*
		We cannot use slices as keys. But we can use arrays.
		The return order of a map is not guaranteed
	*/

	/*
		Use make function for declaring a map
	*/

	myGrades2 := make(map[string]int)
	myGrades2["English"] = 80
	myGrades2["Urdu"] = 70
	myGrades2["Mathematics"] = 100
	myGrades2["Science"] = 92
	fmt.Println(myGrades2)

	/*
		When we try to get a value from a map that does not exist,
		0 is returned. There is no way of knowing if 0 was returned
		because it is the value of the key or because the value
		does not exist in the map by just looking at the value.

		Struct returns another value that we usually name 'ok'. It
		is a boolean value that represents whether the value existed
		in the map or not
	*/
	arabic, ok := myGrades["Arabic"]
	fmt.Println(arabic) // 0
	fmt.Println(ok)     // false

	/*
		Get length of map
	*/
	fmt.Println(len(myGrades2)) // 4
}

/*
User is available globally in other packages,
however, what is inside the struct is not visible.
This is because the struct name starts with a
capital letter and the field names start with lowercase
letters.
If we make one of the pascal case, only that field in the
struct will be visible
*/
type User struct {
	// name  string
	Name   string `required:"true" max:"100"`
	email  string
	age    uint16
	grades map[string]int
}

type Student struct {
	User
	class string
}

func structs() {
	myData := User{
		Name:  "Fassih Haider",
		email: "fassih@gmail.com",
		age:   24,
		grades: map[string]int{
			"English":    90,
			"Urdu":       50,
			"Mathmatics": 100,
		},
	}
	fmt.Println(myData)

	/*
		We can also initialize using positional syntax
		but it is highly discouraged because it makes the code
		less maintainable
	*/
	myData2 := User{
		"Fassih Haider",
		"fassih@gmail.com",
		24,
		map[string]int{
			"English":    90,
			"Urdu":       50,
			"Mathmatics": 100,
		},
	}
	fmt.Println(myData2)

	/*
		we can use anonymous structs for very short lived structs
		that are only used once
	*/
	myName := struct{ name string }{name: "Fassih"}
	fmt.Println(myName)

	/*
		Golang does not support inheritance. It supports composition.
		This is where we can compose types to make new types
	*/
	aStudent := Student{}
	aStudent.age = 24
	aStudent.class = "12"
	aStudent.email = "fassih@gmail.com"
	aStudent.Name = "Fassih"
	fmt.Println(aStudent)

	/*
		If I had to initialize and declare at the sametime, then I would
		have to know the internal structure of the struct that is embedded
	*/
	aStudent2 := Student{
		class: "12",
		User:  User{email: "fassih@gmail.com", age: 24, Name: "Fassih"},
	}
	fmt.Println(aStudent2)

	/*
		We can add tags to struct fields. They do not themselves do something.
		We can get their values and implement our logic. For example 'Name' has
		a tag '`required max:"100"`'
	*/
	t := reflect.TypeOf(User{})
	field, _ := t.FieldByName("Name")
	fmt.Println(field.Tag)
}

func mapsAndStructs() {
	fmt.Println("mapsAndStructs")
	/*
		Maps are REFERENCE types
		Structs are VALUE types
	*/
	maps()
	structs()

}
