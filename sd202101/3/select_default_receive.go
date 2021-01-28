package main

import "fmt"

func main() {

	ch := make(chan int)
	select {
	case <-ch:
		fmt.Println("received")
	default:
	}
}
