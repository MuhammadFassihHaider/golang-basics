package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func deferPanicAndRecovery() {
	fmt.Println("deferPanicAndRecovery")

	// deferring()
	// deferring2()
	// callingAnAPI()
	// callingAnAPI2()
	// deferring3()
	// simplePanic()
	// aSimpleWebServer()
	panicAndDefer()
}

func deferring() {
	/*
		deferred statements are called at the very end
		just before the function ends
	*/
	fmt.Println("=============================")
	fmt.Println("start")
	defer fmt.Println("middle")
	fmt.Println("end")
}

func deferring2() {
	/*
		Defering is LIFO
	*/
	fmt.Println("=============================")
	defer fmt.Println("start")
	defer fmt.Println("middle")
	defer fmt.Println("end")
}

func callingAnAPI() {
	response, err := http.Get("https://www.google.com/robots.txt")

	// checks if there was en error or not
	if err != nil {
		log.Fatal(err)
	}
	// frees up resources after the function execution ends
	defer response.Body.Close()
	readBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s", readBytes)
}

func callingAnAPI2() {
	response, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	readBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", readBytes)
}

func deferring3() {
	a := "Start"
	// it will print "Start" because defer takes the argument eagerly
	defer fmt.Println(a)
	a = "End"
}

func simplePanic() {
	fmt.Println("Start")
	panic("Something horrible happened")
	fmt.Println("End")
}

func aSimpleWebServer() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello"))
	})
	err := http.ListenAndServe(":8080", nil)
	// if we try to re-run the server in another terminal,
	// the server will panic
	if err != nil {
		panic(err.Error())
	}
}

func panicAndDefer() {
	/*
		panic is run after the deferred statements are run
	*/
	fmt.Println("Start")
	defer fmt.Println("Middle")
	panic("Something horrible happened")
	fmt.Println("End")

	/*
		Output:
		Start
		Middle
		Panic
	*/
}
