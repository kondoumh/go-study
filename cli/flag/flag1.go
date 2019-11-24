package main

import (
	"fmt"
	"flag"
)

func main() {
	s := flag.String("s", "Hello, world!", "String help message")
	flag.Parse()
	fmt.Println(*s)
}