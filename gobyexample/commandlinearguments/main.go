package main

import (
	"fmt"
	"os"
)

// go run main.go a b c d
func main() {
	argWithProg := os.Args
	ansWithoutProg := os.Args[1:]
	arg := os.Args[3]
	fmt.Println(argWithProg)
	fmt.Println(ansWithoutProg)
	fmt.Println(arg)
}
