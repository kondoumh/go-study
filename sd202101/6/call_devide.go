package main

import (
	"fmt"
	"os"
	"errors"
)

func divide(x, y int) (int, error) {
	if y == 0 {
		return 0, errors.New("divide by zero")
	}
	return x / y, nil
}

func main() {
	result, err := divide(3, 0)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err.Error())
		os.Exit(1)
	}
	fmt.Printf("result: %d\n", result)
}
