package main

import (
	"log"
	"time"
)

func main() {
	doneCh := make(chan struct{})
	for i := 0; i < 10; i++ {
		i := i
		go do(i, doneCh)
	}
	close(doneCh)
	time.Sleep(300 * time.Millisecond)
}

func do(n int, doneCh <-chan struct{}) {
	for {
		select {
		case <-doneCh:
			log.Printf("finished %d", n)
			return
		default:
			time.Sleep(100 * time.Millisecond)
		}
	}
}
