package main

import "fmt"

type Crier interface {
	Cry() string
}

type ParrotFunc func() string

func (p ParrotFunc) Cry() string {
	return p()
}

func main() {
	var c Crier = ParrotFunc(func() string {
		return "Squawk"
	})
	fmt.Println(c.Cry())
}
