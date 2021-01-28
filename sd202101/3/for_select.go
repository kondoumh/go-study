package main

import (
	"fmt"
	"time"
)

func main() {
	doneCh := make(chan int)
	go func(doneCh chan<- int) {
		time.Sleep(5 * time.Second)
		doneCh <- 1
	}(doneCh)
	for {
		select {
		case <-time.Tick(1 * time.Second):
			fmt.Println("waiting...")
		case <-doneCh:
			return
		}
	}
}
