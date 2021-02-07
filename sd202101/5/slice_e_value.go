package main

import "fmt"

type T struct {
	Number int
	Text   string
}

func main() {
	c := []T{}
	c = append(c, T{})
	v := c[0]
	v.Number = 1
	fmt.Println(c[0].Number) // 0
	c[0].Number = 1
	fmt.Println(c[0].Number) // 1
}
