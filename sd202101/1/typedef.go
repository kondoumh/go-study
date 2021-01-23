package main

import (
	"fmt"
	"time"
)

func main() {
	type MyDuration time.Duration
	d := MyDuration(100)
	fmt.Printf("%T\n", d)
	td := time.Duration(d)
	md := 100 * d
	fmt.Printf("td: %T, md: %T\n", td, md)
}
