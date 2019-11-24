package main

import (
	"fmt"
	"flag"
)

func main() {
	s := flag.String("s", "hello, workd!", "String help message")
	i := flag.Int("i", 1, "Int help message")
	b := flag.Bool("b", false, "Bool help message")

	flag.Parse()

	fmt.Println("str: ", *s)
	fmt.Println("int: ", *i)
	fmt.Println("bool: ", *b)
}
