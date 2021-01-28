package main

import "fmt"

func main() {

	ch := make(chan int)
	select {
	case ch <- 100:
		fmt.Println("sent")
	default:
	}
}
