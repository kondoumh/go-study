package main

import "fmt"

type T struct {
	Number int
	Text   string
}

func main() {
	c := map[int]*T{0: &T}
	c[0].Number = 1
	fmt.Println(c[0].Number)
}
