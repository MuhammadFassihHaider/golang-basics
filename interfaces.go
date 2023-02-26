package main

import (
	"fmt"
	"math"
)

func interfaces() {

	fmt.Println("interfaces")
	// incrementorExample()
	// greeterExample()
	areaAndCircumferenceExample()
}

// >>>>>>>>>>> Area, Cirucmference example >>>>>>>>>>
type shape interface {
	area() float64
	cirucmference() float64
}
type circle struct {
	radius float64
}
type square struct {
	side float64
}

func (c circle) area() float64 {
	r2 := math.Pow(c.radius, 2)
	return math.Pi * r2
}
func (s square) area() float64 {
	return math.Pow(s.side, 2)
}

func (c circle) cirucmference() float64 {
	return 2 * math.Pi * c.radius
}
func (s square) cirucmference() float64 {
	return s.side * 4
}

func printShapeDetails(s shape) {
	fmt.Println("Area: ", s.area())
	fmt.Println("Circumference: ", s.cirucmference())
}

func areaAndCircumferenceExample() {
	var myCircle shape = circle{radius: 5}
	var mySquare shape = square{side: 5}

	printShapeDetails(myCircle)
	printShapeDetails(mySquare)
}

// <<<<<<<<<<< Area, Cirucmference example <<<<<<<<<<

// >>>>>>>>>>> Greeter example >>>>>>>>>>
type Greeter interface {
	greet()
}

type MyGreeting struct {
	name     string
	greeting string
}

func (g MyGreeting) greet() {
	fmt.Println(g.greeting, ", ", g.name)
}
func greeterExample() {
	var sayHello Greeter = MyGreeting{name: "Fassih", greeting: "Hello"}
	sayHello.greet()
}

// <<<<<<<<<<< Greeter example <<<<<<<<<<

// >>>>>>>>>>> Increment on 'int' example >>>>>>>>>>
/*
	1.Interfaces describe the behaviour and not the data.
	2.They are a collection of methods.
	3.Interfaces create abtract types, structs create concrete types
	4.Because interface types are abstract, they cannot be initiated
	like structs
	5.Interfaces in Golang are implemented implicitly
	6.For interfaces that have only one method, they will have the same
	name as the method + er at the end. For example 'Incrementer' has a
	method 'Increment'
*/
func incrementorExample() {
	num := MyInt(0)
	num.Increment()
	num.Increment()
	num.Increment()
	fmt.Println(num)
}

type Incrementer interface {
	Increment() int
}

type MyInt int

// implicit implementation of the 'Incrementer interface'
func (ic *MyInt) Increment() int {
	(*ic)++
	return int(*ic)
}

// <<<<<<<<<<< Increment on 'int' example <<<<<<<<<<
