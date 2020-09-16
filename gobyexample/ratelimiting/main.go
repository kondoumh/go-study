package main

import (
	"fmt"
	"time"
)

func main() {
	request := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		request <- i
	}
	close(request)
	limiter := time.Tick(200 * time.Millisecond)
	for req := range request {
		<-limiter
		fmt.Println("request", req, time.Now())
	}
	burstyLiimiter := make(chan time.Time, 3)

	for i := 0; i < 3; i++ {
		burstyLiimiter <- time.Now()
	}

	go func() {
		for t := range time.Tick(200 * time.Millisecond) {
			burstyLiimiter <- t
		}
	}()

	burstyRequests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		burstyRequests <- i
	}
	close(burstyRequests)
	for req := range burstyRequests {
		<-burstyLiimiter
		fmt.Println("request", req, time.Now())
	}
}
