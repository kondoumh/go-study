package main

import "fmt"

type Crier interface {
	Cry() string
}

type Duck struct{}

func (d Duck) Cry() string {
	return "Quack"
}

func main() {
	var c Crier = Duck{}
	fmt.Println(c.Cry())
}
