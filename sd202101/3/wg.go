package main

import (
	"sync"
	"time"
	"log"
)

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			do(n)
		}(i)
	}
	wg.Wait()
}

func do(n int) {
	time.Sleep(1 * time.Second)
	log.Printf("%d called", n)
}
