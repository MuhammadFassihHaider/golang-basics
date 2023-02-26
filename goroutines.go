package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

func goRoutines() {
	fmt.Println("goroutines")
	goRoutineExample1()
}

func goRoutineExample1() {
	wg := &sync.WaitGroup{}
	mutex := &sync.Mutex{}
	websiteList := []string{
		"https://www.digitalocean.com/",
		"https://golang.org/",
		"https://www.docker.com/",
		"https://sla2ck.com/", // err endpoint
		"https://www.cloudflare.com/",
		"https://www.patreon.com/",
		"https://www.tiobe.com/",
		"https://www.soundcloud.com/",
		"https://www.sendgrid.com/",
		"https://www.coinbase.com/",
	}

	endpointWithStatus := make(map[string]int)

	start := time.Now()

	for _, website := range websiteList {
		wg.Add(1)
		/*
			At the point of writing this code,
			without using goroutines it took ~ 6.4
			With Goroutines, it took ~ 1.5s.
			Ofc it depends on the endpoints which are the
			bottleneck
		*/
		go getStatusCode(website, endpointWithStatus, wg, mutex)
	}

	wg.Wait()

	elapsed := time.Since(start)
	fmt.Println(endpointWithStatus)
	fmt.Println("time taken:", elapsed)
}

func getStatusCode(endpoint string, endpointWithStatus map[string]int, wg *sync.WaitGroup, mutex *sync.Mutex) {
	defer wg.Done()
	defer recoverFromError()

	response, error := http.Get(endpoint)
	if error != nil {
		fmt.Println(error)
		endpointWithStatus[endpoint] = 400
	}
	defer response.Body.Close()

	statusCode := response.StatusCode

	fmt.Printf("Got %d for endpoint %s \n", statusCode, endpoint)

	mutex.Lock()
	endpointWithStatus[endpoint] = statusCode
	mutex.Unlock()
}

func recoverFromError() {
	if err := recover(); err != nil {
		fmt.Println(err)
	}
}
