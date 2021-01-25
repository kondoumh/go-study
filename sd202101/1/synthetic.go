package main

import "fmt"

type Crier interface {
	Cry() string
}

type Footstepper interface {
	FootSteps() string
}

type CryFootstepper interface {
	Crier
	Footstepper
}

type Person struct{}

func (p *Person) Cry() string {
	return "Hi"
}
func (p *Person) FootSteps() string {
	return "Pitapat"
}

type PartyPeaple struct {
	Person
}

func (p *PartyPeaple) Cry() string {
	return "Sup?"
}

func main() {
	var cf CryFootstepper

	cf = &Person{}
	fmt.Println(cf.Cry(), cf.FootSteps())

	cf = &PartyPeaple{}
	fmt.Println(cf.Cry(), cf.FootSteps())
}
